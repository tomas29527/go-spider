module go-spider

replace (
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01 => github.com/golang/net v0.0.0-20180218175443-cbe0f9307d01
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/astaxie/beego v1.11.1
	github.com/jmoiron/sqlx v1.2.0 // indirect
)
