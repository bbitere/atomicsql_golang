package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



type ConfigFile struct {
	Templ_GoLangOrmFile        *TemplateItem
	Templ_GoLangModelFile      *TemplateItem
	Templ_GoLang_SchemaDefItem *TemplateItem
	Templ_GoLang_SchemaDefItem_Col *TemplateItem
	Templ_GoLang_ForeignKey    *TemplateItem

	OutputDBContextFile  string
	Models_Extension     string
	ModelsOutputDir      string
	BaseModelName        string
	ImportPackageModels  string
	ImportPackageOrm     string
	PackageGenSql        string
	ConnectionString     string
	SqlLang              string
	Delimeter            string
	DirJsons             string
}

func (c *ConfigFile) ParseConfigFile(pathFile string) {
	file, err := os.Open(pathFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}

		token := getToken(&line)
		if token == "$" {
			token1 := getToken(&line)
			token2 := getToken(&line)
			if token2 == "=" {
				propertyValue := c.getPropertyValue(strings.TrimSpace(line))
				c.setupSLineProperty(token1, propertyValue)
			}
		} else if token == "@" {
			var content strings.Builder
			for scanner.Scan() {
				line2 := strings.TrimSpace(scanner.Text())
				if line2 == "@#@" {
					break
				}
				content.WriteString(line2)
				content.WriteString("\n")
			}

			token1 := getToken(&line)
			c.setupMLineProperty(token1, content.String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	c.checkProps()
}

func (c *ConfigFile) getPropertyValue(val string) string {
	retValue := ""
	prevIdx := 0
	for idx := 0; idx < len(val); {
		idx = strings_Index(val, "%", idx)
		if idx >= 0 {
			retValue += val[prevIdx:idx]

			idx = idx + 1
			idx2 := strings_Index(val, "%", idx)
			strProp := val[idx:idx2]

			v := os.Getenv(strProp)
			if v != ""  {
				retValue += v
			} else {
				retValue += fmt.Sprintf("%%%s%%", strProp)
			}
			prevIdx = idx2 + 1
			idx = prevIdx
		} else {
			break
		}
	}

	retValue += val[prevIdx:]
	return retValue
}

func (c *ConfigFile) setupSLineProperty(token1, content string) {
	switch token1 {
	case "OutputDBContextFile":
		c.OutputDBContextFile = content
	case "Models_Extension":
		c.Models_Extension = content
	case "ModelsOutputDir":
		c.ModelsOutputDir = content
	case "BaseModelName":
		c.BaseModelName = content
	case "FullName_PackageModels":
		c.ImportPackageModels = content
	case "FullName_PackageOrm":
		c.ImportPackageOrm = content
	case "FullName_PackageGenSql":
		c.PackageGenSql = content
	case "ConnectionString":
		c.ConnectionString = content
	case "SqlLang":
		c.SqlLang = content
	case "DELIMETER":
		c.Delimeter = content
	case "DirJsons":
		c.DirJsons = content
	default:
		fmt.Printf("Not identified token %s in single line property\n", token1)
	}
}

func (c *ConfigFile) setupMLineProperty(token1, content string) {
	switch token1 {
	case "Templ_GoLangOrmFile":
		c.Templ_GoLangOrmFile = NewTemplateItem(token1, content)
	case "Templ_GoLangModelFile":
		c.Templ_GoLangModelFile = NewTemplateItem(token1, content)
	case "Templ_GoLang_SchemaDefItem":
		c.Templ_GoLang_SchemaDefItem = NewTemplateItem(token1, content)
	case "Templ_GoLang_SchemaDefItem_Col":
		c.Templ_GoLang_SchemaDefItem_Col = NewTemplateItem(token1, content)
	case "Templ_GoLang_ForeignKey":
		c.Templ_GoLang_ForeignKey = NewTemplateItem(token1, content)
	default:
		fmt.Printf("Not identified token %s in multiline declaration\n", token1)
	}
}

func (c *ConfigFile) checkProps() {
	retError := c.checkAllProps()
	if retError != "" {
		fmt.Println(retError)
	}
}

func (c *ConfigFile) checkAllProps() string {
	if c.OutputDBContextFile == "" {
		return "Missing prop OutputDBContextFile"
	}

	if c.Models_Extension == "" {
		return "Missing prop Models_Extension"
	}

	if c.ModelsOutputDir == "" {
		return "Missing prop ModelsOutputDir"
	}

	if c.BaseModelName == "" {
		return "Missing prop BaseModelName"
	}

	if c.ImportPackageModels == "" {
		return "Missing prop ImportPackageModels"
	}

	if c.ImportPackageOrm == "" {
		return "Missing prop ImportPackageOrm"
	}

	if c.PackageGenSql == "" {
		return "Missing prop PackageGenSql"
	}

	if c.ConnectionString == "" {
		return "Missing prop ConnectionString"
	}

	if c.SqlLang == "" {
		return "Missing prop SqlLang"
	}

	if c.Delimeter == "" {
		return "Missing prop Delimeter"
	}

	if c.DirJsons == ""  {
		return "Missing prop DirJsons"
	}

	if c.Templ_GoLangOrmFile == nil {
		return "Missing prop Templ_GoLangOrmFile"
	}

	if c.Templ_GoLangModelFile == nil {
		return "Missing prop Templ_GoLangModelFile"
	}

	if c.Templ_GoLang_SchemaDefItem == nil {
		return "Missing prop Templ_GoLang_SchemaDefItem"
	}

	if c.Templ_GoLang_SchemaDefItem_Col == nil {
		return "Missing prop Templ_GoLang_SchemaDefItem_Col"
	}

	if c.Templ_GoLang_ForeignKey == nil {
		return "Missing prop Templ_GoLang_ForeignKey"
	}

	return ""
}

func getToken(line *string) string {
	i := 0
	for ; i < len(*line); i++ {
		ch := (*line)[i]
		if ch == '\t' || ch == '\n' || ch == '\r' || ch == ' ' {
		} else {
			break
		}
	}

	content := ""
	for ; i < len(*line); i++ {
		ch := (*line)[i]
		if ch == '\t' || ch == '\n' || ch == '\r' || ch == ' ' {
			*line = (*line)[i:]
			return content
		}
		content += string(ch)
	}
	*line = (*line)[i:]
	return content
}

func NewTemplateItem(name, text string) *TemplateItem {
	return &TemplateItem{Name: name, Text: strings.ReplaceAll(text, "\r\n", "\n")}
}

func (c *ConfigFile) UseTemplate(original string, template *TemplateItem, dict map[string]string) string {
	if template == nil {
		return original
	}
	return template.ConvertTemplate(dict)
}




type TemplateItem struct {
	Name string
	Text string
}

func (t *TemplateItem) ConvertTemplate(dict map[string]string) string {
	text := t.Text
	templateName := t.Name

	for key, value := range dict {
		val := fmt.Sprintf("@@{%s}", key)
		text = strings.ReplaceAll(text, val, value)
	}

	idx := strings.Index(text, "@@{")
	if idx >= 0 {
		idx2 := strings_Index( text, "}", idx);
		item := text[idx : idx2+1-idx]

		fmt.Printf("Error: tag %s is still present in %s\n", item, templateName)
	}

	return text
}


