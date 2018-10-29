<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

use Google\Protobuf\Internal\Message;
{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Class.Named }} extends Message
{
{{ range .Fields }}
    /**
     * {{ .Anno }}
     * @var {{ .Type }}
     */
    private ${{ .Name }} = {{ .Default }};
{{ end }}

    /**
     * {{ .Name }} constructor.
     * @param array $init
     */
    public function __construct(array $init = [])
    {
        parent::__construct();
    }

{{ range .Fields }}
    /**
     * {{ .Anno }}
     * @return {{ .Type }}
     */
    public function get{{ .Name | Titled }}() : {{ .Type }}
    {
        return $this->{{ .Name }};
    }

    /**
     * {{ .Anno }}
     * @param {{ .Type }} {{ if eq .Repeated true }}...{{ end }}$var
     * @return self
     */
    public function set{{ .Name | Titled }}({{ .Type }} {{ if eq .Repeated true }}...{{ end }}$var) : self
    {
        $this->{{ .Name }} = $var;
        return $this;
    }
{{ end }}
}
