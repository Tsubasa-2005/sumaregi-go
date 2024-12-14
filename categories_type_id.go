package sumaregi

// GetCategoryOpts represents options for GetCategory.
type GetCategoryOpts struct {
	Fields []string `url:"fields,omitempty"` // 検索パラメータ (カンマ区切り指定可)。レスポンス項目を指定可能。一部ネスト項目は指定不可。
}

// GetCategoryResponse represents the response of GetCategory.
type GetCategoryResponse struct {
	CategoryID         string `json:"categoryId"`
	CategoryCode       string `json:"categoryCode"`
	CategoryName       string `json:"categoryName"`
	CategoryAbbr       string `json:"categoryAbbr"`
	DisplaySequence    string `json:"displaySequence"`
	DisplayFlag        string `json:"displayFlag"`
	TaxDivision        string `json:"taxDivision"`
	PointNotApplicable string `json:"pointNotApplicable"`
	TaxFreeDivision    string `json:"taxFreeDivision"`
	ReduceTaxID        string `json:"reduceTaxId"`
	Color              string `json:"color"`
	CategoryGroupID    string `json:"categoryGroupId"`
	ParentCategoryID   string `json:"parentCategoryId"`
	Level              string `json:"level"`
	Tag                string `json:"tag"`
	InsDateTime        string `json:"insDateTime"`
	UpdDateTime        string `json:"updDateTime"`
}
