package {{ .PackageName }};

import javax.persistence.*;
import java.time.LocalDate;

@Entity
@Table(name = "{{ .TableName }}")
public class {{ .EntityName }} {

{{- range .Columns }}
    {{ if .IsPrimaryKey }}
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY) // Ou outra estratégia conforme necessário
    {{ end }}
    @Column(name = "{{ .Name }}"{{ if .IsNullable }}, nullable = true{{ else }}, nullable = false{{ end }}{{ if .IsUnique }}, unique = true{{ end }})
    private {{ .Type }} {{ .Name }};
{{- end }}

    // Getters e Setters

{{- range .Columns }}
    public {{ .Type }} get{{ .Name}}() {
        return {{ .Name }};
    }

    public void set{{ .Name}}({{ .Type }} {{ .Name }}) {
        this.{{ .Name }} = {{ .Name }};
    }
{{- end }}

}
