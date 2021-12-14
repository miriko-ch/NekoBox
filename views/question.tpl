{{template "template/header.tpl" .}}
<link rel="shortcut icon"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200"/>
<link rel="apple-touch-icon"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200"/>
<link rel="icon" sizes="192x192"
      href="{{.userContent.Avatar}}?x-oss-process=image/auto-orient,1/quality,q_70/sharpen,200/resize,limit_0,m_fill,w_200,h_200">
<div>
    <div class="uk-card uk-card-default">
        <div class="uk-card-header">
            <div class="uk-text-left uk-text-small uk-text-muted">{{date .questionContent.CreatedAt "Y-m-d H:i:s"}}</div>
            <h4 class="uk-text-center uk-margin-top uk-margin-bottom">{{.questionContent.Content}}</h4>
        </div>
        {{if ne .questionContent.Answer ""}}
            <div class="uk-card-body">
                <p class="uk-text-small">{{answerFormat .questionContent.Answer}}</p>
                <p class="uk-text-small uk-text-right uk-text-muted">-来自@{{.userContent.Name}}的回答</p>
            </div>
        {{end}}
        <div class="uk-card-footer">
            {{if ne .error ""}}
                <div class="uk-alert-danger" uk-alert>
                    <a class="uk-alert-close" uk-close></a>
                    <p>{{.error}}</p>
                </div>
            {{end}}
            {{if eq .isLogin true}}
                {{if and (eq .user.PageID .pageContent.ID) (eq .questionContent.Answer "") }}
                    <h5 class="uk-text-center">回答问题</h5>
                    {{if and (eq .user.PageID .pageContent.ID) (eq .questionContent.Answer "") }}
                        <form class="uk-float-right"
                              method="post"
                              action="/delete/{{ .pageContent.Domain }}/{{ .questionContent.ID }}">
                            {{ .xsrfdata }}
                            <button class="uk-button uk-button-danger uk-button-small">删除提问</button>
                        </form>
                    {{end}}
                    <form method="post">
                        {{ .xsrfdata }}
                        <div class="uk-margin uk-text-center">
                            <textarea name="answer" class="uk-textarea" rows="3" maxlength="500"
                                      placeholder="在此处撰写你的回答..."></textarea>
                        </div>
                        <div class="uk-margin uk-text-center">
                            <button type="submit" class="uk-button uk-button-primary">回答</button>
                        </div>
                    </form>
                {{end}}
            {{end}}
            <h5 class="uk-text-center">再问点别的问题？</h5>
            <form method="post" action="/_/{{.pageContent.Domain}}" id="form">
                {{ .xsrfdata }}
                <div class="uk-margin uk-text-center">
                    <textarea name="content" class="uk-textarea" rows="3" placeholder="在此处撰写你的问题..."
                              maxlength="1000"></textarea>
                </div>
                <div class="uk-margin uk-text-center">
                    <button type="submit" class="uk-button uk-button-primary g-recaptcha"
                            data-sitekey="{{.recaptcha}}" data-callback="onSubmit">发送提问
                    </button>
                </div>
            </form>

            <hr class="uk-divider-icon">
            <p class="uk-text-left uk-text-muted uk-text-small">@{{ .userContent.Name }}以前回答过的问题</p>
            {{range $index, $elem := .questionsContent}}
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
</div>
{{template "template/footer.tpl" .}}