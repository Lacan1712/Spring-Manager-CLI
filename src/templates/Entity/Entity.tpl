package {{.PackageName}};

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "{{.EntityName}}")
public class {{.EntityName}} {
    public {{.EntityName}}() {

    }
}
