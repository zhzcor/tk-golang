package docx_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
	"testing"
	"tkserver/pkg/docconv"
)

//go:embed 第三方题库题目及试卷导入/管理学与人力资源管理1.docx
var temp []byte

type SetTkQuestion struct {
	Type             int                `json:"type" form:"type" binding:"required,min=1,max=5"`       //题目类型，1：单选题，2：多选题，3：判断题，4：简答题，5：材料题
	QuestionStem     string             `json:"question_stem" form:"question_stem" binding:"required"` //题干(html)
	Options          []TkQuestionOption `json:"options" form:"options"`                                //选项
	KnowledgePoint   []int              `json:"knowledge_points" form:"knowledge_points"`
	Difficulty       int                `json:"difficulty" form:"difficulty" binding:"required,min=1,max=5"` //难易度，1：易，2：较易，3：较难，4：难，5：一般
	Desc             string             `json:"desc" form:"desc"`                                            //解析(html)
	MaterialQuestion []SetTkQuestion    `json:"material_question" form:"material_question"`                  //材料题
}
type TkQuestionOption struct {
	OptionContent string `json:"option_content" form:"option_content"` //选项内容(html)
	OptionName    string `json:"option_name" form:"option_name"`       //选项名称 （string）
	IsRight       int    `json:"is_right" form:"is_right"`             //是否正确 1：否，2：是"
}

func TestParse(t *testing.T) {
	f := bytes.NewReader(temp)

	res, media, err := docconv.ConvertDocx2(f)
	if err != nil {
		t.Fatal(err)
	}

	// 用 <p></p> 切割
	stages := strings.Split(res, "<p></p>")
	fmt.Println(stages)

	type Node struct {
		Items      []string
		childItems []Node
	}

	var doc []SetTkQuestion
	var isQuestionArray bool // 是否是材料题
	var rawItems []*Node
	jumpStage := -1
	for si, stage := range stages {
		// 解析html
		item := &Node{}
		var questionListStartKeyWords = []string{
			"【材料题开始】",
			"[材料题开始]",
		}
		var questionListEndKeyWords = []string{
			"【材料题结束】",
			"[材料题结束]",
		}

		dec := xml.NewDecoder(strings.NewReader(stage))
		for {
			tt, err := dec.Token()
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Fatal(err)

			}
			switch v := tt.(type) {
			case xml.StartElement:
				if v.Name.Local == "file" {
					// 文件处理
					if m, o := media["media"][v.Attr[0].Value]; o {
						// 默认都是图
						// 可选上传oss
						m = "ssss"
						item.Items = append(item.Items, `<img style="max-width:100%;width:100%;" src="`+m+`" />`)
					}

				}
			case xml.CharData:
				if strings.TrimSpace(string(v)) != "" {
					item.Items = append(item.Items, string(v))
				}
			}
		}
		//fmt.Println(len(rawItems), rawItems, rawItems[0])
		if len(item.Items) == 0 {
			continue
		}

		for _, word := range questionListStartKeyWords {
			if strings.Index(item.Items[0], word) != -1 {
				isQuestionArray = true
				jumpStage = si
				titleItems := ""
				for ti, tv := range item.Items {
					if ti > 0 {
						titleItems = titleItems + tv + "\n"
					}
				}
				// 新增一个节点过去
				rawItems = append(rawItems, &Node{
					Items:      []string{titleItems},
					childItems: nil,
				})
			}
		}

		for _, word := range questionListEndKeyWords {
			if strings.Index(item.Items[len(item.Items)-1], word) != -1 {
				isQuestionArray = false
				jumpStage = si
			}
		}

		if isQuestionArray {
			// 直接怼最后一个rawItems的childnode
			if rawItems != nil {
				parentItem := rawItems[len(rawItems)-1]
				if si != jumpStage {
					parentItem.childItems = append(parentItem.childItems, *item)
				}
			}
		} else {
			if jumpStage == si {
				if rawItems != nil {
					parentItem := rawItems[len(rawItems)-1]
					item.Items = item.Items[:len(item.Items)-1]
					parentItem.childItems = append(parentItem.childItems, *item)
				}
			} else {
				rawItems = append(rawItems, item)
			}
		}
	}

	for _, item := range rawItems {
		if item.childItems == nil {
			quest, err := ParseQuestion(item.Items)
			if err != nil {
				t.Fatal(err)
			}
			doc = append(doc, *quest)
		} else {
			parentQuest := SetTkQuestion{Type: 5, QuestionStem: item.Items[0]}
			for _, childItem := range item.childItems {
				subQuest, err := ParseQuestion(childItem.Items)
				if err != nil {
					t.Fatal(err)
				}
				parentQuest.MaterialQuestion = append(parentQuest.MaterialQuestion, *subQuest)
			}
			doc = append(doc, parentQuest)
		}

	}
	d, _ := json.Marshal(doc)
	fmt.Println(string(d))

}

func ParseQuestion(rawItems []string) (*SetTkQuestion, error) {
	// 题型判断
	// 填空题
	packfill := regexp.MustCompile(`\[\[(.*?)\]\]`)

	// 找到答案
	answerKeyWords := []string{
		"答案：",
		"答案:",
		"参考答案：",
		"参考答案:",
		"正确答案：",
		"正确答案:",
		"【答案】",
		"[答案]",
		"【正确答案】",
		"[正确答案]",
		"【参考答案】",
		"[参考答案]",
	}
	// 判断题依据条件
	judgeKeyWords := []string{
		"(正确)",
		"(错误)",
		"（正确）",
		"（错误）",
	}
	// 解析 依据条件
	descKeyWords := []string{
		"【解析】",
		"[解析]",
		"【答案解析】",
		"[答案解析]",
	}
	//材料题结束
	questionListEndKeyWords := []string{
		"【材料题结束】",
		"[材料题结束]",
	}
	questionItem := SetTkQuestion{}
	var rightOptions []string
	/********* 选择，判断题检测 ********/
	jumpLine := 100
	beforeContent := ""
	answerWord := ""
	descWord := ""
	for ri, item := range rawItems {
		var hasRule bool
		ii := packfill.FindAllString(item, -1)
		if len(ii) > 0 {
			questionItem.Type = 6 //判断 填空题
		}
		//fmt.Println(item)

		//图片解析
		im, err := regexp.MatchString(`<img src="`, string(item))
		if err != nil {
			return nil, err
		}
		if im {
			switch beforeContent {
			case "":
				break
			case "answer":
				rightOptions = append(rightOptions, item)
				continue
			case "desc":
				questionItem.Desc = questionItem.Desc + string(item)
				continue
			default:
				for qi, qv := range questionItem.Options {
					if qv.OptionName == beforeContent {
						questionItem.Options[qi].OptionContent = qv.OptionContent + string(item)
						break
					}
				}
				continue
			}
		}

		// 是否是选项
		itemRaw := []rune(item)
		m, err := regexp.MatchString(`[a-zA-Z]`, string(itemRaw[0]))
		if err != nil {
			return nil, err
		}
		if m && ri > 0 && ri < jumpLine {
			// 如果是判断题
			var isRight = 1
			content := string(itemRaw[2:])
			for _, word := range judgeKeyWords {
				if len(itemRaw) > 4 {
					search := itemRaw[:len(itemRaw)-4]
					if strings.Index(item, string(search)) != -1 {
						questionItem.Type = 3
						answerSearch := itemRaw[len(itemRaw)-4:]
						if strings.Index(string(answerSearch), "正确") != -1 {
							isRight = 2
						}
						content = strings.ReplaceAll(content, word, "")
					}
				}
			}

			// 是选项，就添加进去
			optionItem := TkQuestionOption{
				OptionContent: content,
				OptionName:    string(itemRaw[0]),
				IsRight:       isRight,
			}
			beforeContent = string(itemRaw[0])
			questionItem.Options = append(questionItem.Options, optionItem)
			hasRule = true
		}

		// 答案判断
		for wi, word := range answerKeyWords {

			if strings.Index(item, word) != -1 {
				answerWord = word
				jumpLine = wi + 1

			}
		}

		// 解析判断
		for _, word := range descKeyWords {
			if strings.Index(item, word) != -1 {
				descWord = word
				answerWord = ""
			}
		}
		//材料题结束（防止材料题结束录入解析中）
		for _, word := range questionListEndKeyWords {
			if strings.Index(item, word) != -1 {
				descWord = ""
			}
		}

		//答案处理
		if answerWord != "" {
			// 找到答案了，检查是否是 多选题
			// 剔除掉标记
			answerRaw := strings.ReplaceAll(item, answerWord, "")
			answerOptions := []rune(answerRaw)
			beforeContent = "answer"

			for _, option := range answerOptions {
				rightOptions = append(rightOptions, string(option))
			}
			rightOptions = append(rightOptions, "</br>")
			fmt.Println(len(rightOptions), "---", rightOptions)
			hasRule = true
			//answerWord = ""
		}

		//解析处理
		if descWord != "" {
			// 解析
			descRaw := strings.ReplaceAll(item, descWord, "")
			questionItem.Desc += "<p>" + descRaw + "</p>"
			beforeContent = "desc"
			hasRule = true
		}

		// 其它东西直接扔题干
		if !hasRule {
			questionItem.QuestionStem += item
		}

	}
	// 题目类型判断
	if len(questionItem.Options) > 0 {
		if len(rightOptions) == 2 {
			questionItem.Type = 1
		} else if len(rightOptions) > 2 {
			questionItem.Type = 2
		}
	}
	// 处理正确答案
	var newOptions []TkQuestionOption
	for _, option := range questionItem.Options {
		for _, rightOption := range rightOptions {
			if option.OptionName == rightOption {
				option.IsRight = 2
			}
		}
		newOptions = append(newOptions, option)
	}
	questionItem.Options = newOptions

	if questionItem.Type == 0 {
		questionItem.Type = 4
		// 答案放选项里
		questionItem.Options = []TkQuestionOption{
			{
				OptionContent: strings.Join(rightOptions, ""),
			},
		}
	}
	return &questionItem, nil
}
