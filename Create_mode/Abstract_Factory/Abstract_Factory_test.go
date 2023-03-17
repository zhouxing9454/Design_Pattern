package Abstract_Factory

func Save(saveArticle SaveArticle) {
	saveArticle.CreateProse().SaveProse()
	saveArticle.CreateAncientPoetry().SaveAncientPoetry()
}

func ExampleSaveRedis() {
	var factory SaveArticle
	factory = &SaveRedis{}
	Save(factory)
	// Output:
	// Redis Save Prose
	// Redis Save Ancient Poetry
}

func ExampleSaveMysql() {
	var factory SaveArticle
	factory = &SaveMysql{}
	Save(factory)
	// Output:
	// Mysql Save Prose
	// Mysql Save Ancient Poetry
}
