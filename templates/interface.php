<?php
# source: {{ .Meta.Source }}

namespace {{ .Contract.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

interface {{ .Contract.Named }}
{
{{ range .Methods }}
    /**
     * {{ .Anno }}
     * @var {{ .Input }} $request
     * @return {{ .Output }}
     */
    public function {{ .Name }}({{ .Input }} $request);
{{ end }}
}
