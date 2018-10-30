<?php
# source: {{ .Meta.Source }}

namespace {{ .Client.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Client.Named }} extends \Carno\RPC\Client implements \{{ .Contract }}
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
