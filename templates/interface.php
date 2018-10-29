<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

interface {{ .Class.Named }}
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
