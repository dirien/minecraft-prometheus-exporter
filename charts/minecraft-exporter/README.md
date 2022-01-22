## Motivation

Currently, I am working on a Helm chart for my Minecraft Prometheus exporter. Nothing special on this task. But now I want to put as much automation into it, as possible.

## Auto Docs With helm-docs

The helm-docs tool auto-generates documentation from helm charts into markdown files. That is very cool, as we don't need to take care to keep the doc in sync with the `values.yaml`

In the chart directory I created a `README.md.tpl`, where I created the template for the final `README.md` file.

```markdown
# Minecraft Exporter for Prometheus

![Version: {{ .Version }}](https://img.shields.io/badge/Version-{{ .Version | replace "-" "--" }}-informational?style=for-the-badge)
{{ if .Type }}![Type: {{ .Type }}](https://img.shields.io/badge/Type-{{ .Type }}-informational?style=for-the-badge) {{ end }}
{{ if .AppVersion }}![AppVersion: {{ .AppVersion }}](https://img.shields.io/badge/AppVersion-{{ .AppVersion | replace "-" "--" }}-informational?style=for-the-badge) {{ end }}

![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)
![Minecraft](https://img.shields.io/badge/Minecraft-62B47A?style=for-the-badge&logo=Minecraft&logoColor=white)
![Docker](https://img.shields.io/badge/docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Alpine Linux 3.15.0](https://img.shields.io/badge/alpine_linux_3.15.0-0D597F?style=for-the-badge&logo=alpine-linux&logoColor=white)

## Description

{{ template "chart.description" . }}

## Usage
todo

{{ template "chart.valuesSection" . }}

{{ template "chart.homepageLine" . }}

{{ template "chart.sourcesSection" . }}

{{ template "chart.maintainersSection" . }}
```

`helm-docs` has plenty of predefined templates for you to use, and you also use the variables directly. Plus mix your own static text in. That's very handy in the usage section.

To get the values table rendered with proper default values, you need to add `# -- description`

```yaml
image:
  # -- The docker image repository to use
  repository: ghcr.io/dirien/minecraft-exporter
  # -- The docker image tag to use
  tag: ''
  # -- The docker image pull policy
  pullPolicy: IfNotPresent
```
You can add `# -- @default  defaut text`

e.g:

```yaml
image:
  # -- The docker image repository to use
  repository: ghcr.io/dirien/minecraft-exporter
  # -- The docker image tag to use
  # @default Chart version
  tag: ''
  # -- The docker image pull policy
  pullPolicy: IfNotPresent
```

Now you have to option to create a `pre-commit` hook, so everytime you change a specified file in your `.pre-commit-hooks.yaml` the `README.md` file gets rendered

#### pre-commit

I am going to use [pre-commit](https://pre-commit.com/). Check the docs on how to install it.

I created a `git-hook` folder with the `helm-docs.sh` with following content:

```bash
#!/usr/bin/env bash

set -e

if ! command -v helm-docs > /dev/null 2>&1; then
    echo "Please install helm-docs to run the pre-commit hook! https://github.com/norwoodj/helm-docs#installation"
    exit 1
fi

helm-docs "${@}"
```

The `.pre-commit-config.yaml` contains the location of the `.pre-commit-hooks.yaml` on the remote repository.

```yaml
repos:
  - repo: https://github.com/dirien/minecraft-prometheus-exporter
    rev: v0.10.0
    hooks:
      - id: helm-docs
        args:
          - --template-files=README.md.gotmpl
```
The content of the `.pre-commit-hooks.yaml` defines on which files, the `git-hook/helm-docs.sh` should be called. And of course the relation between the `id` of the `.pre-commit-config.yaml`

```yaml
- id: helm-docs
  args: []
  description: Uses 'helm-docs' to create documentation from the Helm chart's 'values.yaml' file, and inserts the result into a corresponding 'README.md' file.
  entry: git-hook/helm-docs.sh
  files: (README\.md\.gotmpl|(Chart|requirements|values)\.yaml)$
  language: script
  name: Helm Docs
  require_serial: true
```

During the start, you will not have a remote location with your `.pre-commit-hooks.yaml` so you can test your hook via this command:

```bash
pre-commit install
pre-commit try-repo . helm-docs --verbose --all-files
```
Otherwise, you can just call:

```bash
pre-commit install
pre-commit install-hooks
```

>The only downside of this approach is, If someone is contributing to your `helm` chart, >you need to rely on that has the `pre-commit` hook cli installed.

#### CI

So I will add the generation of the helm-docs to my CI pipeline.

```yaml
      - name: Run helm-docs
        run: |
          cd /tmp
          wget https://github.com/norwoodj/helm-docs/releases/download/v$HELM_DOCS_VERSION/helm-docs_$HELM_DOCS_VERSION_Linux_x86_64.tar.gz
          tar -xvf helm-docs_$HELM_DOCS_VERSION_Linux_x86_64.tar.gz
          sudo mv helm-docs /usr/local/sbin
          helm-docs -t chart/README.md.tpl -o README.md
          ls -la chart/
          cat chart/README.md
```

You can find more information here -> https://github.com/norwoodj/helm-docs

## Chart Testing & Linting

`ct` is the tool for testing Helm charts. It automatically detects charts changed against the target branch.

I created a config file called `ct-lint.yaml` under the  `.github/configs` with following content:

```yaml
remote: origin
target-branch: main
chart-dirs:
  - charts
helm-extra-args: "--timeout 600s"
validate-chart-schema: true
validate-chart-values: true
validate-maintainers: true
validate-yaml: true
exclude-deprecated: true
excluded-charts: []
```
Defining some properties. On addition, I created `lintconf.yaml`under the same folder.

```yaml
---
rules:
  braces:
    min-spaces-inside: 0
    max-spaces-inside: 0
    min-spaces-inside-empty: -1
    max-spaces-inside-empty: -1
.....  snip .....
  without
    check-multi-line-strings: false
  key-duplicates: enable
  line-length: disable # Lines can be any length
  new-line-at-end-of-file: enable
  new-lines:
    type: unix
  trailing-spaces: enable
  truthy:
    level: warning
```
And then you can run `ct lint  --config .github/configs/ct-lint.yaml --lint-conf .github/configs/lintconf.yaml` to execute the linting

As `ct` needs  some tools like `Helm, Git (2.17.0 or later), Yamllint, Yamale, Kubectl` it is good to use the provided `ct` container.

```bash
docker container run --rm -v $(pwd):/workspace -ti quay.io/helmpack/chart-testing:v3.5.0
```

If you have a local kubernetes cluster you can call:

```bash
ct install --config ./.github/configs/ct-lint.yaml
```

It will install the chart and if everything works fine instantly deletes it.

```bash
Installing charts...

------------------------------------------------------------------------------------------------------------------------
 Charts to be processed:
------------------------------------------------------------------------------------------------------------------------
 minecraft-exporter => (version: "0.1.0", path: "charts/minecraft-exporter")
------------------------------------------------------------------------------------------------------------------------

Installing chart 'minecraft-exporter => (version: "0.1.0", path: "charts/minecraft-exporter")'...
Creating namespace 'minecraft-exporter-hwe920t1yv'...
namespace/minecraft-exporter-hwe920t1yv created
NAME: minecraft-exporter-hwe920t1yv

....

========================================================================================================================
Deleting release 'minecraft-exporter-hwe920t1yv'...
release "minecraft-exporter-hwe920t1yv" uninstalled
W0122 21:56:16.272288   69990 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
Deleting namespace 'minecraft-exporter-hwe920t1yv'...
namespace "minecraft-exporter-hwe920t1yv" deleted
Namespace 'minecraft-exporter-hwe920t1yv' terminated.
------------------------------------------------------------------------------------------------------------------------
 ✔︎ minecraft-exporter => (version: "0.1.0", path: "charts/minecraft-exporter")
------------------------------------------------------------------------------------------------------------------------
All charts installed successfully
```


Now, we can create a GitHub action to run during a PR and test and lint our Helm chart.

```yaml
---
name: Linting and Testing
on: pull_request
jobs:
  chart-test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0

      - name: Set up Helm
        uses: azure/setup-helm@v1
        with:
          version: v3.6.3

      - name: Set up python
        uses: actions/setup-python@v2
        with:
          python-version: 3.7

      - name: Setup Chart Linting
        id: lint
        uses: helm/chart-testing-action@v2.2.0

      - name: List changed charts
        id: list-changed
        run: |
          ## If executed with debug this won't work anymore.
          changed=$(ct --config ./.github/configs/ct-lint.yaml list-changed)
          charts=$(echo "$changed" | tr '\n' ' ' | xargs)
          if [[ -n "$changed" ]]; then
            echo "::set-output name=changed::true"
            echo "::set-output name=changed_charts::$charts"
          fi

      - name: Run chart-testing (lint)
        run: ct lint --debug --config ./.github/configs/ct-lint.yaml --lint-conf ./.github/configs/lintconf.yaml

      - name: Create kind cluster
        uses: helm/kind-action@v1.2.0
        if: steps.list-changed.outputs.changed == 'true'

      - name: Run chart-testing (install)
        run: ct install --config ./.github/configs/ct-lint.yaml
        if: steps.list-changed.outputs.changed == 'true'
```

We use the predefined GitHub action for installing the helm/chart-testing CLI tool.

On top as you can see, we start a `Kind` cluster to install and test the `helm` chart.

You can find more information here -> https://github.com/helm/chart-testing

## Chart Releaser

The last part, is the quickest. We use `cr` as tool here. `cr` is a tool designed to help us to release a helm chart and connect a GitHub releases with the `index.yaml` and upload this to the GitHub pages.

We just need to add in our `chart-publish` following task at as an additional task:

```yaml
     - name: Run chart-releaser
        uses: helm/chart-releaser-action@v1.2.1
        env:
          CR_TOKEN: "${{ secrets.GITHUB_TOKEN }}"
```
That's it.

You can find more information here -> https://github.com/helm/chart-releaser

## Resources

- https://github.com/norwoodj/helm-docs
- https://github.com/helm/chart-testing
- https://github.com/helm/chart-releaser
- https://github.com/marketplace/actions/helm-chart-testing