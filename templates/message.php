<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Class.Named }} extends \Google\Protobuf\Internal\Message
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
        {{ .Meta | MDInit }}
        parent::__construct();
    }

{{ range .Fields }}
    /**
     * {{ .Anno }}
     * @return {{ .Type }}{{ if eq .Repeated true }}[]{{ end }}
     */
    public function get{{ .Name | Titled }}() : {{ if eq .Repeated true }}array{{ else }}{{ .Type }}{{ end }}
    {
        {{ if eq .Repeated true }}
        $list = [];
        foreach ($this->{{ .Name }}->getIterator() as $item) {
            $list[] = $item;
        }
        return $list;
        {{ else }}
        return $this->{{ .Name }};
        {{ end }}
    }

    /**
     * {{ .Anno }}
     * @param {{ .Type }} {{ if eq .Repeated true }}...{{ end }}$var
     * @return self
     */
    public function set{{ .Name | Titled }}({{ .Type }} {{ if eq .Repeated true }}...{{ end }}$var) : self
    {
        $this->{{ .Name }} = {{ if .Mapped }}{{ $.GPBUtil }}::checkMapField($var, {{ .Mapped.Key }}, {{ .Mapped.Val }}){{ else }}$var{{ end }};
        return $this;
    }
{{ end }}
}
