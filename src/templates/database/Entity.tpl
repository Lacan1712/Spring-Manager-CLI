package {{ .PackageName }};

public class {{ .EntityName }} {

{{- range .Columns }}
    private {{ .Type }} {{ .Name }};
{{- end }}

    // Getters e Setters

{{- range .Columns }}
    public {{ .Type }} get{{ .Name }}() {
        return {{ .Name }};
    }

    public void set{{ .Name }}({{ .Type }} {{ .Name }}) {
        this.{{ .Name }} = {{ .Name }};
    }
{{- end }}

}
