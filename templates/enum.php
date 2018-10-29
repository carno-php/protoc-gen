<?php
# source: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

class {{ .Class.Named }}
{
{{ range .Values }}
    /**
     * {{ .Anno }}
     */
    public const {{ .Key }} = {{ .Val }};
{{ end }}
}
