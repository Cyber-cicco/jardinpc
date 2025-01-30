// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/views/components"
	"github.com/Cyber-cicco/jardin-pc/internal/views/evenements"
)

func Home(evts []*model.Evenement) templ.Component {
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
		templ_7745c5c3_Err = components.Head("Jardin Pollen et Collemboles").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"flex flex-col w-full bg-blue-50\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.NavBar([]components.NavLink{
			{Name: "Accueil", Link: "/home"},
			{Name: "Événements", Link: "/events"},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main id=\"main\" class=\"flex flex-col overflow-auto w-full \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = InnerHome(evts).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func InnerHome(evts []*model.Evenement) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section hx-history=\"false\" class=\"flex flex-col h-screen w-full items-center gap-2 justify-center zoom-in bg-main\"><h1 data-fade-in class=\"font-bold font-poppins text-center px-3 text-white text-4xl lg:text-6xl\">Bienvenu au jardin Pollen et Collemboles ! </h1><h2 class=\"italic text-white fade-right-1 text-xl lg:text-3xl\">Le jardin partagé du quartier Richolets</h2></section><div class=\"flex flex-col gap-10 w-full\"><section class=\"flex flex-row w-full\"><div class=\"flex flex-col p-2 lg:flex-row justify-center\"><div class=\"flex p-10 justify-end\"><img alt=\"les habitants au travail !\" data-fade-in height=\"500\" width=\"500\" class=\"rounded-xl border-2 border-black\" src=\"/img/travail.jpg\"></div><div class=\"flex flex-col w-full p-5 lg:p-2 lg:w-2/3 gap-5 bg-white border-green-400 border-2 shadow-green-400 text-blue-950 shadow-md bg-opacity-35 items-center justify-center rounded-lg\"><h3 data-fade-right class=\"text-center font-poppins font-bold text-4xl\">Qui sommes-nous ?</h3><div class=\"flex justify-center w-full\"><article data-fade-in class=\"flex flex-col w-full lg:w-2/3 gap-2 text-lg text-justify\"><p>Habitants du quartier Richolets à Saint-Herblain, nous avons pris l'initiative de reprendre la main sur le jardin partagé, abandonné depuis quelques temps.</p><p>Notre but ? En faire un lieu d'échange d'animation et de convivialité, un espace vert au milieu des tours en béton. Nous avons donc décider de remettre en état le jardin, pour ensuite pouvoir y mettre en place des animations (gouters, chantiers collectif, plantations, etc.)</p></article></div></div></div></section><section class=\"flex flex-row-reverse w-full\"><div class=\"flex flex-col p-2 lg:flex-row-reverse justify-center\"><div class=\"flex p-10 justify-end\"><img alt=\"photo de maps pour localiser le jardin : 47.21407477563655, -1.6101690670270175\" data-fade-in height=\"500\" width=\"500\" class=\"rounded-xl border-2 border-black\" src=\"/img/localisation.png\"></div><div class=\"flex flex-col w-full p-5 lg:p-2 lg:w-2/3 gap-5 bg-white border-green-400 border-2 shadow-green-400 text-blue-950 shadow-md bg-opacity-35 items-center justify-center rounded-lg\"><h3 data-fade-left class=\"text-center font-poppins font-bold text-4xl\">Comment nous rejoindre ?</h3><div class=\"flex justify-center w-full\"><article data-fade-in class=\"flex flex-col w-full lg:w-2/3 gap-2 text-lg text-justify\"><p>Pour nous rejoindre, rien de plus simple !</p><p>Il vous suffit simplement de venir au jardin à une date où un événement est prévu. <bold>Et pas besoin d'avoir la main verte !</bold> Les personnes déjà présentes s'assurent de la transmission des compétences. Regardez simplement les événements à venir sur ce site, ou <bold>suivez notre page instagram / facebook pour ne rien manquer</bold>.</p></article></div></div></div></section></div><div class=\"p-10\"></div><div class=\"flex justify-center w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = evenements.EvenementsSection(evts).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"py-5\"></div><footer class=\"flex p-5 flex-col bg-blue-600 text-white h-[500px]\"><h4 class=\"italic text-3xl\">Nos réseaux :</h4><div class=\"réseaux p-5 flex gap-5\"><a target=\"_blank\" href=\"https://www.instagram.com/jardin_pollen_et_collemboles/\"><svg alt=\"instagram icon by https://icons8.com\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 40 40\" width=\"80px\" height=\"80px\"><path fill=\"#8585cc\" d=\"M30.5,38.5c4.418,0,8-3.582,8-8v-21c0-4.418-3.582-8-8-8h-21c-4.418,0-8,3.582-8,8v21 c0,4.418,3.582,8,8,8H30.5z\"></path><path fill=\"#8d8dd8\" d=\"M3.4,4.331C2.217,5.726,1.5,7.528,1.5,9.5v21c0,4.418,3.582,8,8,8h21c4.418,0,8-3.582,8-8v-21 c0-0.503-0.052-0.992-0.141-1.469C32.135,4.22,24.832,2,17,2C12.229,2,7.657,2.832,3.4,4.331z\"></path><path fill=\"#bd82f4\" d=\"M1.505,9.404C1.504,9.437,1.5,9.468,1.5,9.5v21c0,4.418,3.582,8,8,8h21c4.418,0,8-3.582,8-8V12.897 C32.439,8.56,25.021,6,17,6C11.465,6,6.22,7.226,1.505,9.404z\"></path><path fill=\"#ed73f4\" d=\"M1.5,13.88V30.5c0,4.418,3.582,8,8,8h21c4.418,0,8-3.582,8-8V17.981C32.724,13.013,25.217,10,17,10 C11.394,10,6.124,11.414,1.5,13.88z\"></path><path fill=\"#f97dcd\" d=\"M17,14c-5.705,0-11.014,1.664-15.5,4.509V30.5c0,4.418,3.582,8,8,8h21c4.418,0,8-3.582,8-8v-6.935 C33.194,17.698,25.534,14,17,14z\"></path><path fill=\"#fc9c95\" d=\"M17,18c-5.861,0-11.237,2.033-15.5,5.411V30.5c0,4.418,3.582,8,8,8h21c4.418,0,8-3.582,8-8v-0.238 C34.143,22.925,26.152,18,17,18z\"></path><path fill=\"#ffac99\" d=\"M17,22c-6.145,0-11.66,2.651-15.5,6.859V30.5c0,4.418,3.582,8,8,8h21c2.465,0,4.668-1.117,6.136-2.87 C33.648,27.674,25.999,22,17,22z\"></path><path fill=\"#ffc49c\" d=\"M30.5,38.5c0.957,0,1.87-0.177,2.721-0.485C31.087,31.065,24.649,26,17,26 c-6.186,0-11.592,3.309-14.566,8.248C3.778,36.777,6.437,38.5,9.5,38.5H30.5z\"></path><path fill=\"#ffde8d\" d=\"M17,30c-5.137,0-9.573,2.984-11.684,7.309C6.535,38.06,7.964,38.5,9.5,38.5h19.683 C27.35,33.542,22.595,30,17,30z\"></path><path fill=\"#fff69f\" d=\"M17,34c-3.319,0-6.193,1.813-7.753,4.487C9.332,38.49,9.415,38.5,9.5,38.5h15.26 C23.203,35.818,20.324,34,17,34z\"></path><path fill=\"#8b75a1\" d=\"M31,2c3.86,0,7,3.14,7,7v22c0,3.86-3.14,7-7,7H9c-3.86,0-7-3.14-7-7V9c0-3.86,3.14-7,7-7H31 M31,1H9 C4.582,1,1,4.582,1,9v22c0,4.418,3.582,8,8,8h22c4.418,0,8-3.582,8-8V9C39,4.582,35.418,1,31,1L31,1z\"></path><path fill=\"#fff\" d=\"M27.5 11A1.5 1.5 0 1 0 27.5 14A1.5 1.5 0 1 0 27.5 11Z\"></path><path fill=\"none\" stroke=\"#fff\" stroke-miterlimit=\"10\" stroke-width=\"2\" d=\"M20 14A6 6 0 1 0 20 26A6 6 0 1 0 20 14Z\"></path><path fill=\"none\" stroke=\"#fff\" stroke-miterlimit=\"10\" stroke-width=\"2\" d=\"M33,14.5c0-4.142-3.358-7.5-7.5-7.5 c-2.176,0-8.824,0-11,0C10.358,7,7,10.358,7,14.5c0,2.176,0,8.824,0,11c0,4.142,3.358,7.5,7.5,7.5c2.176,0,8.824,0,11,0 c4.142,0,7.5-3.358,7.5-7.5C33,23.324,33,16.676,33,14.5z\"></path></svg></a> <a target=\"_blank\" href=\"https://www.facebook.com/groups/636894522144645/\"><svg alt=\"facebook icon by https://icons8.com\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 48 48\" width=\"80px\" height=\"80px\"><path fill=\"#039be5\" d=\"M24 5A19 19 0 1 0 24 43A19 19 0 1 0 24 5Z\"></path><path fill=\"#fff\" d=\"M26.572,29.036h4.917l0.772-4.995h-5.69v-2.73c0-2.075,0.678-3.915,2.619-3.915h3.119v-4.359c-0.548-0.074-1.707-0.236-3.897-0.236c-4.573,0-7.254,2.415-7.254,7.917v3.323h-4.701v4.995h4.701v13.729C22.089,42.905,23.032,43,24,43c0.875,0,1.729-0.08,2.572-0.194V29.036z\"></path></svg></a></div><a href=\"/conditions\" hx-get=\"/conditions\" hx-push-url hx-target=\"#main\" class=\"italic text-3xl\">Consulter les conditions d'utilisation</a></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
