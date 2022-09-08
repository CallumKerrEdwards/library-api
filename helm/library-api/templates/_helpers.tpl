{{/*
Expand the name of the chart.
*/}}
{{- define "library-api.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "library-api.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "library-api.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "library-api.labels" -}}
helm.sh/chart: {{ include "library-api.chart" . }}
{{ include "library-api.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "library-api.selectorLabels" -}}
app.kubernetes.io/name: {{ include "library-api.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "library-api.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "library-api.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create the name of the database secret to use
*/}}
{{- define "library-api.databaseSecretName" -}}
{{- if .Values.mongodb.secret.create }}
{{- default (printf "%s-db" (include "library-api.fullname" .)) .Values.mongodb.secret.name }}
{{- else }}
{{- default "default" .Values.mongodb.secret.name }}
{{- end }}
{{- end }}


{{- define "library-cms.mongodbHost" -}}
{{- if .Values.mongodb.fromRelease.enable }}{{ .Release.Name }}{{ .Values.mongodb.fromRelease.hostSuffix }}
{{- else }}
{{- default "" .Values.mongodb.host }}
{{- end }}
{{- end }}
