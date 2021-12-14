{{template "template/header.tpl" .}}
<link rel="shortcut icon"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200"/>
<link rel="apple-touch-icon"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200"/>
<link rel="icon" sizes="192x192"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200">
<div class="uk-card uk-card-default uk-text-center">
    <div class="uk-height-medium uk-flex uk-flex-center uk-flex-middle uk-background-cover uk-light"
         data-src="{{ .pageContent.Background }}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/format,png"
         uk-img>
        <div class="uk-card-body">
            <img class="uk-border-circle uk-box-shadow-large"
                 src="{{ .userContent.Avatar }}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200/format,png"
                 width="100" height="100">
            <h3>{{ .userContent.Name }}</h3>
            <p>{{ .pageContent.Intro }}</p>
        </div>
    </div>
</div>

<div>
    <div class="uk-card uk-card-default uk-card-small uk-card-body">
        <p class="uk-text-center uk-text-small">谁都可以以匿名的形式提问</p>
        <form method="post" id="form">
            {{ .xsrfdata }}
            {{if ne .error ""}}
                <div class="uk-alert-danger" uk-alert>
                    <a class="uk-alert-close" uk-close></a>
                    <p>{{.error}}</p>
                </div>
            {{end}}
            {{if ne .success ""}}
                <div class="uk-alert-success" uk-alert>
                    <a class="uk-alert-close" uk-close></a>
                    <p>{{.success}}</p>
                </div>
            {{end}}
            <div class="uk-margin uk-text-center">
                <textarea name="content" class="uk-textarea" rows="3" placeholder="在此处撰写你的问题..."
                          maxlength="1000">{{.questionDraft}}</textarea>
            </div>
            <div class="uk-margin uk-text-center">
                <button type="submit" class="uk-button uk-button-primary g-recaptcha" data-sitekey="{{.recaptcha}}"
                        data-callback="onSubmit">发送提问
                </button>
            </div>
        </form>
        <hr class="uk-divider-icon">
        <p class="uk-text-left uk-text-muted uk-text-small">@{{ .userContent.Name }}以前回答过的问题</p>
        {{range $index, $elem := .questionContent}}
            {{ if ne $elem.Answer ""}}
                <div>
                    <hr>
                    <a class="uk-button uk-button-default uk-button-small uk-float-right"
                       href="/_/{{$.pageContent.Domain}}/{{$elem.ID}}">查看回答</a>
                    <div class="uk-text-left uk-text-small uk-text-muted">{{date $elem.CreatedAt "Y-m-d H:i:s"}}</div>
                    <p class="uk-text-small">{{$elem.Content}}</p>
                </div>
            {{end}}
        {{end}}
    </div>
</div>
{{template "template/footer.tpl" .}}