cd ..
~/go/bin/templ generate
npx tailwindcss -i tailwind.base.css -o resources/static/css/tailwind.css
cd internal
go run main.go

