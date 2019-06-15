# terraformer-plan-splitter
## list

```bash
$ terraformer-plan-modifier list <planfile>
```

## split

```bash
$ terraformer-plan-modifier split <patternfile> <planfile>
```

### patternfile

```yaml
defaultname: default
mappings:
  - name: example
    patterns:
      - ^(www\.)?example\.com$
```

Patterns support [RE2 syntax](https://github.com/google/re2/wiki/Syntax).
