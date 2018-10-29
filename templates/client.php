<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

use Carno\RPC\Client;
{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Class.Named }} extends Client implements API
{
{{ range .Methods }}
    /**
     * {{ .Anno }}
     * @var {{ .Input }} $request
     * @return {{ .Output }}
     */
    public function {{ .Name }}({{ .Input }} $request)
    {
        return $this->request('{{ .Package }}', '{{ .Service }}', '{{ .Name }}', $request, new {{ .Output }});
    }
{{ end }}
}
