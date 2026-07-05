package htmx

import (
	"bufio"
	"context"

	"github.com/catcher3/cegla"
)

var (
	keyGet     = []byte("hx-get")
	keyPost    = []byte("hx-post")
	keyPut     = []byte("hx-put")
	keyDelete  = []byte("hx-delete")
	keyTarget  = []byte("hx-target")
	keySwap    = []byte("hx-swap")
	keyTrigger = []byte("hx-trigger")
	keyConfirm = []byte("hx-confirm")
	keyPrompt  = []byte("hx-prompt")

	keySse = []byte("hx-sse")
	keyWs  = []byte("hx-ws")
)

// Attr встраивает cegla.AttrMarker — как и atr.Attr, доступен в любой
// категории контента (Flow, Phrasing, TableRowContent, ListChild и т.д.),
// поэтому htmx.Post(...) можно положить и в Button{}, и в TR{}, и куда угодно.
type Attr struct {
	cegla.AttrMarker
	key []byte
	val []byte
}

func (Attr) IsAttribute()   {}
func (h Attr) Name() string { return string(h.key) }

func (h Attr) Key() string   { return string(h.key) }
func (h Attr) Value() string { return string(h.val) }

// Render ничего не пишет напрямую — атрибуты обрабатываются в
// cegla.RenderChildren/cegla.RenderVoid при рендере родительского тега.
func (h Attr) Render(ctx context.Context, w *bufio.Writer) error {
	return nil
}

// Внутренний хелпер для сокращения кода
func attr(key, val []byte) Attr {
	return Attr{key: key, val: val}
}

// Запросы

// Get отправляет GET-запрос по указанному URL.
func Get(url string) Attr { return attr(keyGet, []byte(url)) }

// Post отправляет POST-запрос по указанному URL.
func Post(url string) Attr { return attr(keyPost, []byte(url)) }

// Put отправляет PUT-запрос по указанному URL.
func Put(url string) Attr { return attr(keyPut, []byte(url)) }

// Delete отправляет DELETE-запрос по указанному URL.
func Delete(url string) Attr { return attr(keyDelete, []byte(url)) }

// Управление поведением

// Target указывает CSS-селектор элемента, который должен быть обновлен ответом сервера.
func Target(selector string) Attr { return attr(keyTarget, []byte(selector)) }

// Swap определяет способ замены контента (например, "innerHTML", "outerHTML").
func Swap(method string) Attr { return attr(keySwap, []byte(method)) }

// Trigger задает событие, которое инициирует запрос (по умолчанию "click").
func Trigger(event string) Attr { return attr(keyTrigger, []byte(event)) }

// Confirm показывает диалоговое окно подтверждения перед отправкой запроса.
func Confirm(msg string) Attr { return attr(keyConfirm, []byte(msg)) }

// Prompt показывает диалоговое окно с полем ввода перед отправкой запроса.
func Prompt(msg string) Attr { return attr(keyPrompt, []byte(msg)) }

// Server side events

// SseConnect устанавливает соединение с сервером для получения событий (SSE).
func SseConnect(url string) Attr { return attr(keySse, []byte("connect "+url)) }

// SseSwap определяет имя сообщения SSE, при получении которого происходит замена контента.
func SseSwap(msgName string) Attr { return attr(keySse, []byte("swap "+msgName)) }

// Websocket

// WsConnect устанавливает WebSocket-соединение с указанным URL.
func WsConnect(url string) Attr { return attr(keyWs, []byte("connect "+url)) }

// WsSend отправляет сообщение через активное WebSocket-соединение при наступлении события.
func WsSend(event string) Attr { return attr(keyWs, []byte("send "+event)) }

// Request Modifiers

// Include включает значения других элементов (по селектору) в отправляемый запрос.
func Include(selector string) Attr { return attr([]byte("hx-include"), []byte(selector)) }

// Headers добавляет пользовательские HTTP-заголовки к запросу.
func Headers(json string) Attr { return attr([]byte("hx-headers"), []byte(json)) }

// Vals добавляет дополнительные значения (JSON) в параметры запроса.
func Vals(json string) Attr { return attr([]byte("hx-vals"), []byte(json)) }

// Sync синхронизирует запросы (например, отменяет предыдущие или ставит в очередь).
func Sync(rule string) Attr { return attr([]byte("hx-sync"), []byte(rule)) }

// PushUrl обновляет текущий URL в адресной строке браузера без перезагрузки страницы.
func PushUrl(url string) Attr { return attr([]byte("hx-push-url"), []byte(url)) }

// Indicator указывает селектор элемента, который будет получать класс 'htmx-request' во время загрузки.
func Indicator(sel string) Attr { return attr([]byte("hx-indicator"), []byte(sel)) }
