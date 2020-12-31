package i18n_test

import (
	"fmt"
	"golang.org/x/text/feature/plural"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 手把手教你 Go 程序的国际化和本土化 https://zyfdegh.github.io/post/201805-translation-go-i18n/
// A Step-by-Step Guide to Go Internationalization (i18n) & Localization (l10n) https://phrase.com/blog/posts/internationalization-i18n-go/

func init() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")
	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")
	message.SetString(language.Chinese, "How are you?", "你好吗?.")
}

func ExampleI18n() {
	p := message.NewPrinter(language.Chinese)
	p.Printf("%s went to %s.", "彼得", "英格兰")
	fmt.Println()
	p.Printf("%s has been stolen.", "宝石")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")

	// Output:
	// 彼得去了英格兰。
	// 宝石被偷走了。
	// Peter is in England.
	// The Gem has been stolen.
}

func init() {
	message.Set(language.English, "我有 %d 个苹果",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))
	message.Set(language.English, "还剩余 %d 天",
		plural.Selectf(1, "%d",
			"one", "One day left",
			"other", "%[1]d days left",
		))

}

func ExamplePlural() {
	p := message.NewPrinter(language.English)
	p.Printf("我有 %d 个苹果", 1)
	fmt.Println()
	p.Printf("我有 %d 个苹果", 2)
	fmt.Println()
	p.Printf("我有 %d 个苹果", 5)
	fmt.Println()
	p.Printf("还剩余 %d 天", 1)
	fmt.Println()
	p.Printf("还剩余 %d 天", 10)
	fmt.Println()

	// Output:
	// I have an apple
	// I have two apples
	// I have 5 apples
	// One day left
	// 10 days left
}
