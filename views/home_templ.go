// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "andrenormanlang/tarot-go-htmx/common"

func Home(cards []common.Card, selectedCards []common.Card, meanings []string, isShuffling bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Free Tarot Reading</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"/static/css/styles.css\" rel=\"stylesheet\"><link href=\"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css\" rel=\"stylesheet\"></head><body class=\"bg-gray-100 text-center py-2 px-2\"><div class=\"container mx-auto max-w-4xl\"><h1 class=\"text-2xl sm:text-3xl font-bold mb-2\">Rider Waite Smith (Mind, Body & Spirit)</h1><!-- Button to trigger the modal --><button type=\"button\" class=\"text-xl sm:text-2xl btn btn-info w-64 mb-2 smythe-regular\" data-bs-toggle=\"modal\" data-bs-target=\"#infoModal\">Meaning of Card Positions?</button><div class=\"w-full max-w-lg mx-auto mb-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(selectedCards) == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"shuffleButton\" class=\"mb-2 text-xl sm:text-2xl bg-purple-500 hover:bg-purple-600 w-64 text-white px-4 py-2 rounded-lg smythe-regular transition duration-300 ease-in-out\">Shuffle Cards</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" placeholder=\"Enter your question or subject here (optional)\" id=\"tarot-question\" class=\"old-standard-tt-bold w-full p-2 border border-gray-300 rounded mb-2\"></div><!-- Modal for Card Positions --><div class=\"modal fade\" id=\"infoModal\" tabindex=\"-1\" aria-labelledby=\"infoModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog modal-lg\"><div class=\"modal-content\"><div class=\"modal-header\"><h5 class=\"text-2xl modal-title smythe-regular\" id=\"infoModalLabel\">Card Positions in the \"Mind, Body & Spirit\" Tarot Spread</h5><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body px-6 py-4\"><table class=\"w-full border-collapse\"><thead><tr class=\"bg-gray-200 text-gray-700\"><th class=\"text-3xl p-4 text-left smythe-regular\">Position</th><th class=\"text-3xl p-4 text-center smythe-regular\">Meaning</th></tr></thead> <tbody><tr class=\"border-b border-gray-300\"><td class=\"text-2xl p-4 font-semibold smythe-regular\">MIND:</td><td class=\"text-2xl p-4 text-center old-standard-tt-regular text-justify\">The first card represents logic, reasoning, and your rational side. What are you thinking? Or: What is an area of reason you should pay attention to?</td></tr><tr class=\"border-b border-gray-300\"><td class=\"text-2xl p-4 font-semibold smythe-regular\">BODY:</td><td class=\"text-2xl p-4 text-center old-standard-tt-regular text-justify\">The second card represents health, fitness, sexuality, and activity. What are you doing? Or: What is an area of action you should think about?</td></tr><tr class=\"border-b border-gray-300\"><td class=\"text-2xl p-4 font-semibold smythe-regular\">SPIRIT:</td><td class=\"text-2xl p-4 text-center old-standard-tt-regular text-justify\">The third and last card represents emotion, creativity, spirituality, and love. What are you feeling? Or: What should you focus on for spiritual growth?</td></tr></tbody></table></div><div class=\"modal-footer\"><button type=\"button\" class=\"btn btn-secondary\" data-bs-dismiss=\"modal\">Close</button></div></div></div></div><!-- Do Another Reading Button -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(selectedCards) == 3 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"doAnotherReading\" hx-get=\"/reset-reading\" hx-target=\"body\" hx-swap=\"innerHTML\" class=\"mt-4 text-xl sm:text-2xl bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg  smythe-regular transition duration-300 ease-in-out\">Do Another Reading</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Shuffled Cards Container --><div id=\"shuffled-cards\" class=\"relative h-36 sm:h-48 mt-2 flex justify-center items-center overflow-x-auto mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, card := range cards {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card-container\" style=\"--card-index: {index};\"><button hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("/select-card?card=" + card.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 75, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" class=\"bg-transparent\" onclick=\"openModal(&#39;{ card.Name }&#39;, &#39;{ card.MeaningUp }&#39;)\"><img src=\"/images/CardBacks.png\" alt=\"Card Back\" class=\"w-24 sm:w-32 h-auto\"></button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Selected Cards Container --><div id=\"selected-cards\" class=\"mt-2 grid grid-cols-1 sm:grid-cols-3 gap-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if selectedCards != nil {
			for i, card := range selectedCards {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center\"><h3 class=\"text-lg sm:text-xl font-bold mb-2 smythe-regular\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if i == 0 {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Mind (Past)")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else if i == 1 {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Body (Present)")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Spirit (Future)")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><div class=\"flip-card\"><div class=\"flip-card-inner\"><div class=\"flip-card-front\"><img src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("/images/" + card.Image)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 98, Col: 45}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(card.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 98, Col: 63}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full h-auto\"></div><div class=\"flip-card-back flex flex-col justify-between p-2\"><h2 class=\"text-sm sm:text-md font-semibold text-white\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(card.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 101, Col: 78}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"text-xs sm:text-sm text-white flex-grow overflow-y-auto\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(card.MeaningUp)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 102, Col: 94}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><button type=\"button\" class=\"btn btn-primary mt-1 old-standard-tt-regular text-sm sm:text-lg\" hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/card-detail?card=" + card.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 106, Col: 53}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cardModal .modal-content\" data-bs-toggle=\"modal\" data-bs-target=\"#cardModal\">View Full Description</button></div></div></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Modal for Extended Card Description --><div class=\"modal fade\" id=\"cardModal\" tabindex=\"-1\" aria-labelledby=\"cardModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog modal-lg modal-dialog-centered\"><div class=\"modal-content\"><div class=\"modal-header\"><h5 class=\"modal-title text-center\" id=\"cardModalLabel\">Card Title</h5><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body px-6 py-4\"><p class=\"text-lg old-standard-tt-bold mb-2\">Description:</p><p id=\"modal-desc\" class=\"text-base old-standard-tt-regular mb-4\"></p><hr class=\"my-4 border-gray-300\"><p class=\"text-lg old-standard-tt-bold mb-2\">Meaning (Upright):</p><p id=\"modal-meaning\" class=\"text-base old-standard-tt-regular\"></p></div><div class=\"modal-footer\"><button type=\"button\" class=\"btn btn-secondary\" data-bs-dismiss=\"modal\">Close</button></div></div></div></div><!-- External Scripts --><script src=\"https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.7/dist/umd/popper.min.js\"></script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js\"></script><script src=\"/static/script/htmx.min.js\"></script><script src=\"/static/script/script.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
