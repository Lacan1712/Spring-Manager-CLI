package {{.PackageName}}

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/{{.ControllerName}}")
public class {{.ControllerName}} {

    @GetMapping
    public String getAll() {
        return "Hello from {{.ControllerName}}";
    }

    @PostMapping
    public String create() {
        return "Created";
    }
}
