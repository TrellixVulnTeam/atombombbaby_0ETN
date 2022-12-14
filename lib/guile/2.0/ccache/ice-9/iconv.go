GOOF----LE-4-2.0?      ] 3 4   h?      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  ice-9?	g  iconv?	 ?		g  filenameS?	
f  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?	g  importsS?	g  rnrs?	g  bytevectors?	 ?	 ?	g  binary-ports?	 ?	 ?	g  rdelim?	 ?	g  selectS?	g  read-string?	 ?	 ?	 ?	g  exportsS?	g  string->bytevector?	g  bytevector->string?	g  call-with-encoded-output-string?	 ?	g  set-current-module?	  ?	! ?	"g  open-output-string?	#g  get-output-string?	$g  
close-port?	%g  call-with-output-string*?	&g  open-bytevector-output-port?	'g  call-with-output-bytevector*?	(g  error?	)g  set-port-encoding!?	*g  set-port-conversion-strategy!?	+g  string-ci=??	,f  utf-8?	-g  string->utf8?	.g  display?	/g  open-bytevector-input-port?	0g  eof-object??	1f   ?	2g  utf8->string?C 5   h?  ?   ]4	
5 4! >  "  G   "#$  h@     ]45 4 >  "  G  454>  "  G  C             g  proc
		9 g  port		9 g  str		"	9  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	
??			??			??	
		??		 	??	"	 	??	%	!	?? 		9  g  nameg  call-with-output-string*? C%R&$      h@   .  ]4>   G 4 >  "  G  45 4>  "  G  C&      g  proc
		@ g  port		@ g  get-bytevector			@ g  bv		)	@  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	$
??		%	??		%	??		'	??	%	(	??	)	(	??	,	)	?? 		@  g  nameg  call-with-output-bytevector*? C'R(')*  h@   ?   ]4 L>  "  G  L$  4 L>  "  G  "   L  6      ?       g  port
		:  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	8	??		9		??		:		??		;	??	:	<		?? 		:   C+,-%    hX   ?  - . , 3 #  "   O 64 5$  &  456"???"???       ?      g  encoding
		Q g  proc		Q g  conversion-strategy			Q  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	,
??		.	???	)	7	??	)	1	??	*	1	??	0	1	!??	2	1	??	6	1	??	:	2	$??	>	1	??	A	6	??	I	6	?? 		Q	  g  nameg  call-with-encoded-output-string?g  documentationf  yCall PROC on a fresh port.  Encode the resulting string as a
bytevector according to ENCODING, and return the bytevector.? CR(.       h   ?   ]L  6      ?       g  port
		
  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	J	??	
	K		?? 		
   C+,-      hP   ?  - . , 3 #  "   O 645$  &   6"???"???    ?      g  str
		L g  encoding		L g  conversion-strategy			L  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	A
??		B	=??	)	H	??	)	E	??	*	E	??	0	E	!??	2	E	??	6	E	??	:	F	$??	>	E	??	D	G	?? 		L	  g  nameg  string->bytevector?g  documentationf  hEncode STRING according to ENCODING, which should be a string naming
a character encoding, like "utf-8".? CR(/)*$01+,2       h?   ?  - . , 3 #  "  g4 54>  "  G  $  4>  "  G  "   454>  "  G  45$  CC4	
5$  &   6"??z"??v    y      g  bv
	 ? g  encoding	 ? g  conversion-strategy		 ? g  p		! ? g  res		\ ?  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?
	N
??		O	=??		V	??	!	V	??	$	W	??	=	X	??	>	Y	??	V	Z	??	\	Z	??	_	[	
??	q	\	??	{	\	
??	}	]	?? ?	S	?? ?	S	?? ?	S	!?? ?	S	?? ?	S	?? ?	T	$?? ?	S	?? ?	U	?? 	 ?	  g  nameg  bytevector->string?g  documentationf  ?Decode the string represented by BV.  The bytes in the bytevector
will be interpreted according to ENCODING, which should be a string
naming a character encoding, like "utf-8".? CRC  ?       g  m
		,  g  filenamef  S/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/iconv.scm?		
???	
??	$
??~	,
??	o	A
???	N
?? 	?
   C6 