package template

const GetTemplate = `URL: 
{{ .Url }}
Headers: 
{{ .Headers }}
Response: 
{{ .Response }}
`

const PostTemplate = `URL: 
{{ .Url }}
Headers: 
{{ .Headers }}
Params: 
{{ .Params }}
Response: 
{{ .Response }}
`
