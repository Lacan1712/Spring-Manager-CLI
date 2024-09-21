package {{.PackageName}}

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/{{.ClassName}}")
public class {{.ClassName}} {

    @GetMapping
    public String getAll() {
        return "Hello from {{.ClassName}}";
    }

    @PostMapping
    public String create() {
        return "Created";
    }
}
