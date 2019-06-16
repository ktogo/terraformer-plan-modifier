# Terraformer Plan Modifier

CLI tool to modify the [_Plan file_](https://github.com/GoogleCloudPlatform/terraformer#planning) generated from and used for [Terraformer](https://github.com/GoogleCloudPlatform/terraformer).

## List command

`list` command lists all resouces in the given planfile.

```bash
$ terraformer-plan-modifier list <planfile>
```

## Split command

`split` command split planfile based on given mapping file.

```bash
$ terraformer-plan-modifier split --mapping <mapfile> --plan <planfile>
```

### Mapping file

```yaml
defaultname: default
mappings:
  - name: example
    selector: .InstanceState.Attributes.name
    patterns:
      - ^(www\.)?example\.com$
```

#### Selector

Selector specifies which value to compare against the patterns.

Selector syntax supports [text/template syntax](https://golang.org/pkg/text/template/).

Surrounding double braces `{{ ... }}` will be autocompleted if the selector begin with dot (`.`) and doesn't contain any brace characters.

#### Patterns

Patterns tests whether given value specified by the selector matches to the condition.

Pattern syntax supports [RE2 syntax](https://github.com/google/re2/wiki/Syntax).
