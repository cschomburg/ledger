package print

var (
	tplPlainXact = `{{ to_date .Date }}{{ if not .EffectiveDate.IsZero }} = {{ to_date .EffectiveDate }}
{{- end}}{{ if .IsPending }} ! {{end}}
{{- if .IsCleared }} * {{end}}
{{- with .Code }} ({{ .Code }}) {{end -}}
{{ .Description }}{{ with .Note }}{{ .NotePreSpace }}{{.Note}}
{{- end}}

{{- $node := . -}}

{{ range .Postings }}
{{ posting_account_pre_space $node . }}
{{- if .IsPending }}! {{ end}}
{{- if .IsCleared }}* {{ end -}}
{{ .Account }}{{ posting_account_post_space $node . }}
{{- with .BalanceAssertion }}= {{ amount . }}
{{- end}}{{ with .Amount }}{{ amount . }}
{{- end}}{{ with .LotPrice }} { {{- amount . -}} }
{{- end}}{{ if not .LotDate.IsZero }} [{{ to_date . }}]
{{- end}}{{ with .Price }} {{ if .PriceIsForWhole }}@@{{else}}@{{end}} {{ amount . }}
{{- end}}{{ with .BalanceAssertion }} = {{ amount . }}
{{- end}}{{end}}
`
)