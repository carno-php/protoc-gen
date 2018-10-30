<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Class.Named }} extends \Carno\RPC\Client implements {{ .Contracted }}
{
{{ range .Methods }}
    /**
     * {{ .Anno }}
     * @var {{ .Input }} $request
     * @return {{ .Output }}
     */
    public function {{ .Name }}({{ .Input }} $request)
    {
        return $this->request('{{ $.Package }}', '{{ $.Name }}', '{{ .Name }}', $request, new {{ .Output }});
    }
{{ end }}
}
