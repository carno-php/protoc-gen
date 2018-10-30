<?php
# file: {{ .Meta.Source }}

namespace {{ .Class.Namespaced }};

{{ range .CTX.Namespaces }}use {{ . }};{{ "\n" }}{{ end }}

class {{ .Class.Named }}
{
    private static $initialized = false;

    /**
     * {{ .Class.Named }} initialize
     */
    public static function init() : void
    {
        if (self::$initialized) {
            return;
        }

        {{ range .Imports }}
        {{ . }}::init();{{ end }}

        \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool()->internalAddGeneratedFile(hex2bin(
        {{ range .Lines }}
            "{{ . }}" . {{ end }}
            ""
        ));

        self::$initialized = true;
    }
}
