{{- define "name" -}}
{{ .Release.Name }}
{{- end -}}

{{- define "image" -}}
{{ printf "%s:%s" .Values.image.repository (default (printf "v%s" .Chart.AppVersion) .Values.image.tag) }}
{{- end -}}

{{- define "controller" -}}
{{ printf "%s-controller" .Release.Name }}
{{- end -}}

{{- define "node" -}}
{{ .Release.Name }}
{{- end -}}

{{- define "serviceAccount" -}}
{{ default .Release.Name .Values.serviceAcocunt }}
{{- end -}}

{{- define "secret" -}}
{{ default .Release.Name .Values.secretName }}
{{- end -}}

{{- define "namespace" -}}
{{ .Release.Namespace }}
{{- end -}}

{{- define "timestamp" -}}
{{ date "20060102150405" .Release.Time | quote }}
{{- end -}}
