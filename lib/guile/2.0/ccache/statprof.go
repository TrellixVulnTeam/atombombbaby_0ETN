GOOF----LE-4-2.0??      ] ? 4       h?      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  statprof?	 ?	g  filenameS?		f  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?	
g  importsS?	g  srfi?	g  srfi-1?	 ?	 ?	g  system?	g  vm?	 ?	 ?	g  frame?	 ?	 ?	g  program?	 ?	 ?	 ?	g  exportsS?	g  statprof-active??	g  statprof-start?	g  statprof-stop?	g  statprof-reset?	g  statprof-accumulated-time?	 g  statprof-sample-count?	!g  statprof-fold-call-data?	"g  statprof-proc-call-data?	#g  statprof-call-data-name?	$g  statprof-call-data-calls?	%g  statprof-call-data-cum-samples?	&g  statprof-call-data-self-samples?	'g  statprof-call-data->stats?	(g  statprof-stats-proc-name?	)g  statprof-stats-%-time-in-proc?	*g  statprof-stats-cum-secs-in-proc?	+g   statprof-stats-self-secs-in-proc?	,g  statprof-stats-calls?	-g  !statprof-stats-self-secs-per-call?	.g   statprof-stats-cum-secs-per-call?	/g  statprof-display?	0g  statprof-display-anomolies?	1g  statprof-fetch-stacks?	2g  statprof-fetch-call-tree?	3g  with-statprof?	4g  gcprof?	5 !"#$%&'()*+,-./01234 ?	6g  	autoloadsS?	7g  ice-9?	8g  format?	978 ?	:8 ?	;9: ?	<g  set-current-module?	=< ?	>< ?	?g  accumulated-time?	@g  last-start-time?	Ag  sample-count?	Bg  sampling-frequency?	Cg  remaining-prof-time?	Dg  profile-level?	Eg  %count-calls??	Fg  gc-time-taken?	Gg  record-full-stacks??	Hg  stacks?	Ig  procedure-data?	Jg  make-call-data?	Kg  call-data-proc?	Lg  procedure-name?	Mg  call-data-name?	Ng  with-output-to-string?	Og  write?	Pg  call-data-printable?	Qg  call-data-call-count?	Rg  call-data-cum-sample-count?	Sg  call-data-self-sample-count?	Tg  inc-call-data-call-count!?	Ug  inc-call-data-cum-sample-count!?	Vg   inc-call-data-self-sample-count!?	Wg  make-syntax-transformer?	XW ?	YW ?	Zg  accumulate-time?	[g  macro?	\g  $sc-dispatch?	]\ ?	^\ ?	_g  _?	`g  any?	a_`??	bg  syntax->datum?	cb ?	db ?	eg  datum->syntax?	fe ?	ge ?	hg  set!?	ig  +?	je  0.0?	kg  -?	l@ ?	mg  syntax-violation?	nm ?	om ?	pf  -source expression failed to match any pattern?	qg  program??	rg  program-num-free-variables?	sg  program-objcode?	tg  	hashq-ref?	ug  
hashq-set!?	vg  get-call-data?	wg  stack-length?	xg  frame-procedure?	yg  
count-call?	zg  frame-previous?	{g  make-hash-table?	|g  	hash-fold?	}g  and=>?	~g  	stack-ref?	g  sample-stack-procs? ?g  inside-profiler?? ?g  get-internal-run-time? ?g  
make-stack? ?g  profile-signal-handler? ?g  pk? ?g  what!? ?g  set-vm-trace-level!? ?g  the-vm? ?g  vm-trace-level? ?g  	setitimer? ?g  ITIMER_PROF? ?g  assq? ?g  gc-stats? ?g  	add-hook!? ?g  vm-apply-hook? ?g  remove-hook!? ?g  error? ?f  /Can't reset profiler while profiler is running.? ?g  	sigaction? ?g  SIGPROF? ?f  :Can't call statprof-fold-called while profiler is running.? ?e  100.0? ?e  1.0? ?g  max? ?g  stats-sorter? ?g  current-output-port? ?f  No samples recorded.
? ?g  sort? ?f  !~5a ~10a   ~7a ~8a ~8a ~8a  ~8@a
? ?f  %  ? ?f  
cumulative? ?f  self? ?f   ? ?f  total? ?f   ~5a  ~9a  ~8a ~8a ~8a ~8a  ~8@a
? ?f  time? ?f  seconds? ?f  calls? ?f  ms/call? ?f  name? ?f  ~5a ~10a   ~7a  ~8@a
? ?f  %? ?f  ~5a  ~10a  ~7a  ~8@a
? ?g  for-each? ?f  #~6,2f ~9,2f ~9,2f ~7d ~8,2f ~8,2f  ? ?f  ~6,2f ~9,2f ~9,2f  ? ?g  display? ?g  newline? ?f  ---
? ?g  simple-format? ?f  Sample count: ~A
? ?f  *Total time: ~A seconds (~A seconds in GC)
? ?g  internal-time-units-per-second? ?f  ==[~A ~A ~A]
? ?f  Total time: ~A
? ?f  5Can't get accumulated time while profiler is running.? ?g  procedure=?? ?g  map? ?g  lists->trees? ?g  cadr? ?g  find? ?g  	assq-set!? ?g  filter? ?g  identity? ?g  unfold-right? ?g  stack->procedures? ?g  loopS? ???? ?g  hzS? ??	?? ?g  count-calls?S? ??	?? ?g  full-stacks?S? ??	?? ????? ? ?g  inexact->exact? ?g  floor? ?e  	1000000.0? ?f  Invalid macro body? ?g  keyword?? ?g  eq?? ?g  @? ?? ? ?g  lambda? ??	?? ??? ? ?g  after-gc-hook?C 5h?o  ?  ]4	
56;5	 4> >  "  G   ?R@RARBRCR
DRER
FRGRHRIR  h     ]  C        g  proc
		 g  
call-count		 g  cum-sample-count			 g  self-sample-count			  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	?? 			  g  nameg  make-call-data? CJR    h   ?   ] 
?C ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	?? 		  g  nameg  call-data-proc? CKRLK     h   ?   ]4 56   ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	,??	 ?	?? 		  g  nameg  call-data-name? CMRMNOK       h   ?   ] 4L 56   ?       g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?	??	 ?	/??	 ?	(?? 		
   C       h    ?   ]	4 5$  C O 6 ?       g  cd
		 g  t			  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	??		 ?	??	 ?	?? 		  g  nameg  call-data-printable? CPR      h   ?   ] ?C ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	"?? 		  g  nameg  call-data-call-count? CQR     h   ?   ] 	?C?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	(?? 		  g  nameg  call-data-cum-sample-count? CRR       h   ?   ] 	?C?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	)?? 		  g  nameg  call-data-self-sample-count? CSR      h   ?   ]  ???C   ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??		 ?	??	
 ?	??	 ?	?? 		  g  nameg  inc-call-data-call-count!? CTR      h   ?   ] 	 	???C ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	??	 ?	??	 ?	?? 		  g  nameg  inc-call-data-cum-sample-count!? CURh   ?   ] 	 	???C ?       g  cd
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	??	 ?	??	 ?	?? 		  g  nameg   inc-call-data-self-sample-count!? CVR4YZ[^adgh?ijkl        h    ?   ] ??  C     ?       g  	stop-time
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	?? 		   C h   ?   ]	4 5L 4?6?       g  args
		 g  v			  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?	 ?
?? 		   Cop    h(   ?   ]	4 5$   O @ 6 ?       g  y
		' g  tmp		'  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
?? 		'   C5ZRqrstIJu      hp   g  ]4 5?$  "  	4 5
?$   "  4 545$  C4 


54>  "  G  C _      g  proc
		o g  t	
	# g  k	4	o g  t		?	o g  	call-data		T	o  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	??	
 ?	??	
 ?	??	 ?	??	 ?	??	' ?	??	. ?	??	4 ?	??	7 ?	??	? ?	??	K ?	??	T ?	??	W ?	
?? 		o  g  nameg  get-call-data? CvRwGHAxyz{u|Uv h   ?   ]4 56   ?       g  proc
		 g  val		 g  accum			  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?		??	 ?	??	 ?	?? 			   C}vV~        h  @  ])H4 >  "  G  $   ? "   ? "  ?$  ?45$  Z&  K454	5"???4	>  "  G  45$  "  "???45"??|4
>  "  G  445>  "  !G  "  4 
54	5"??+JC       8      g  stack
		 g  hit-count-call?		 g  frame		1 ? g  
procs-seen		1 ? g  self		1 ? g  t		> ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
 ?
??	 ?	??	 ?	??	 ?	??	" ?	??	$ ?	??	+ ?	??	- ?	??	1 ?	??	7 ?	??	8 ?	??	> ?	??	N ?	??	Q	??	R	??	Y	,??	j	??	k	?? ?	?? ?		?? ?	?? ?	?? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ?	??  		  g  nameg  sample-stack-procs? CR?R?D?????E????j@??B h?   ?  ] 
?$  ?45 45$  "  445545$  "  7	$  $4
45 445 5?>  "  G  "   ??? 4

??>  "  G  $  "  145  	$  $4
45 445 5?>  "  G  "   "    C   ?      g  sig
	 ? g  	stop-time	 ? g  t			: g  stack		: ? g  inside-apply-trap?		C ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?

??		??		??		??		??		??		??		??	+	??	/	??	0	#??	7	??	:	??	=	!??	C	??	K!	??	U)	??	V*	??	Y*	'??	^+	+??	a+	;??	g+	+??	h+	'??	m*	??	} ?	?? ?,	?? ?.	?? ?0	?? ?1	?? ?.	?? ?3	?? ?5	$?? ?5	?? ?6	?? ?7	?? ?7	'?? ?8	+?? ?8	;?? ?8	+?? ?8	'?? ?7	?? ?:	?? +	 ?  g  nameg  profile-signal-handler? C?R??j?@}xTv       h   ?   ]4 56   ?       g  proc
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
E	??	G	??	F	?? 		   C 	   h@   ?   ]$  C?45 ?? 44 5>  "  G  45  C    ?       g  frame
		<  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	@	??	 ?	??	B	??	B	??	D	??	D	??	+D	??	4I	??	:I	?? 		<  g  nameg  
count-call? CyRD        h   4  ] 
?C ,      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
M
??	P	?? 		
  g  nameg  statprof-active??g  documentationf  uReturns @code{#t} if @code{statprof-start} has been called more times
than @code{statprof-stop}, @code{#f} otherwise.? CRDC?@?F?F??BE???y?? h?   ?  ]? ?$  ?  $   ?
?$  "   ?
?"   45  445 5? $  4	


 ? ?>  "  G  "  4	


??>  "  G  $   4445 5>  "  G  "   445 445 5?>  "  G  CC    ?      g  rpt
	 ? g  t		4 g  use-rpt?	9 ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
S
??	W	??	W	??	X	??	X	??	Y	??	Z	??	[	+??	[	 ??	[	??	/\	+??	0\	 ??	9Y	??	>]	??	?^	??	E^	??	F`	??	J`	??	K`	(??	Q`	??	R`	??	T_	??	Za	??	[b	??	db	'??	gb	1??	lb	??	yc	?? ?e	?? ?f	?? ?c	?? ?g	?? ?h	?? ?h	?? ?h	&?? ?h	?? ?h	?? ?i	?? ?i	?? ?i	*?? ?i	:?? ?i	*?? ?i	&?? ?i	?? -	 ?
  g  nameg  statprof-start?g  documentationf  Start the profiler.@code{}? CRD?F?F???E??y??C?j?@  h?   ?  ] ? 
?$  z445 5?? 445 445 5?>  "  G  	$   4
445 5>  "  G  "   4



5 ?45 ??  CC     ?      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
m
??	q	??	q	??	
r	??	r	??	u	??	u	??	u	+??	u	??	u	??	u	??	!t	??	"v	??	%v	??	*v	*??	-v	:??	3v	*??	4v	&??	9v	??	Gw	??	Hx	??	Kx	??	Nx	)??	Tx	??	[x	??	h{	"??	t{	??	x ?	??	z|	?? ?|	?? ?}	??  	 ?
  g  nameg  statprof-stop?g  documentationf  Stop the profiler.@code{}? CRD??E?@ABC{IGH???    hx     - . , 3 #  
?$  4>  "  G  "    
  
  ?  	4
 ?5   4>  "  G  C             g  sample-seconds
		q g  sample-microseconds		q g  count-calls?			q g  full-stacks?			q  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?

??	?	??	?	??	?	??	"?	??	'?	??	7?	??	:?	??	=?	??	@?	??	E?	??	G?	??	J?	??	K?	??	T?	??	X?	??	Y?	??	[?	??	\?	?? 		q	  g  nameg  statprof-reset?g  documentationf VReset the statprof sampler interval to @var{sample-seconds} and
@var{sample-microseconds}. If @var{count-calls?} is true, arrange to
instrument procedure calls as well as collecting statistical profiling
data. If @var{full-stacks?} is true, collect all sampled stacks into a
list for later analysis.

Enables traps and debugging as necessary.? CRD??|   h   ?   ]L 6      ?       g  key
		
 g  value		
 g  prior-result			
  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	
?	?? 		
	   CI       h0   O  ]
?$  4>  "  G  "    O 6 G      g  proc
		/ g  init		/  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??	
?	??	?	??	?	??	?	??	/?	?? 		/	  g  nameg  statprof-fold-call-data?g  documentationf 7Fold @var{proc} over the call-data accumulated by statprof. Cannot be
called while statprof is active. @var{proc} should take two arguments,
@code{(@var{call-data} @var{prior-result})}.

Note that a given proc-name may appear multiple times, but if it does,
it represents different functions with the same name.? C!RD??tI        h(   Z  ]
?$  4>  "  G  "    6R      g  proc
		(  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??	
?	??	?	??	?	??	?	??	(?	?? 		(  g  nameg  statprof-proc-call-data?g  documentationf  TReturns the call-data associated with @var{proc}, or @code{#f} if
none is available.? C"RPSR E$??j? h?     ]14 54 54 545 45 45 ?$  4 5"  ???	??	?$  
?$  
"  ?	??"  $  ?	?45?"   C   ?      g  	call-data
	 ? g  	proc-name		 ? g  self-samples		 ? g  cum-samples		 ? g  all-samples		" ? g  secs-per-sample		/ ? g  	num-calls		C ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??		?	??	?	??	?	??	?	??	?	??	?	??	"?	??	%?	??	*?	??	/?	??	/?	??	7?	??	8?	'??	C?	??	L?	??	N?	,??	O?	??	T?	??	V?	+??	W?	??	\?	??	^?	,??	_?	??	g?	??	j?	??	o?	??	q?	*??	z?	??	|?	9??	}?	?? ??	?? ??	?? ??	?? ??	?? ??	?? ??	?? ??	?? (	 ?  g  nameg  statprof-call-data->stats?g  documentationf  0Returns an object of type @code{statprof-stats}.? C'R h   ?   ] 
?C ?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	)?? 		  g  nameg  statprof-stats-proc-name? C(R      h   ?   ] ?C ?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	.?? 		  g  nameg  statprof-stats-%-time-in-proc? C)R h   ?   ] 	?C?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	0?? 		  g  nameg  statprof-stats-cum-secs-in-proc? C*R       h   ?   ] 	?C?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	1?? 		  g  nameg   statprof-stats-self-secs-in-proc? C+R      h   ?   ] 	?C?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	%?? 		  g  nameg  statprof-stats-calls? C,R  h   ?   ] 	?C?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	2?? 		  g  nameg  !statprof-stats-self-secs-per-call? C-R     h   ?   ] 	?C?       g  stats
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	1?? 		  g  nameg   statprof-stats-cum-secs-per-call? C.R+*        h8   '  ]
4 545?
?$  4 545?"  
?C           g  x
		3 g  y		3 g  diff			3  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??	
?	??	?	??	?	??	?		??	?	??	?	??	#?	??	*?		??	1?	?? 		3	  g  nameg  stats-sorter? C?R? 8?!'     h   ?   ]4 5?C   ?       g  data
		 g  prior-value		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	 ??	?	?? 			   C??E????????????????E8?)*+,-.??(?  h?   ?  ]$  J4M 4 54 54 54 5?4 5??4	 5?>  "  G  "  )4M 
4 54 54 5>  "  G  44 5M >  "  G  M 6      y      g  stats
	 ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??		?	??	?	??	?	??	?	??	?	??	%?	??	/?	??	6?	??	:?	??	A?	??	F?	??	S?	??	Y?	??	Z?	??	a 	??	h	??	s?	??	|	??		?? ?	?? ?	?? 	 ?  g  nameg  display-stats-line? C?????F?? $ h(  ?  -  1  3  H J (  45 K "   45 
?$  J 64545	$  H4J 
>	  "  G  4J >	  "  G  "  84J >  "  G  4J >  "  G  4 O >  "  G  4J >  "  G  445 >  "  G  4 5 !"?#?6   ?      g  port
	% g  
stats-list	9% g  sorted-stats		D%  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??	?	??	?	??	?	??	$?	??	)?	??	/?	??	1?	??	2?	??	7?	??	9?	??	9?	??	<?	??	D?	??	L	??	M	??	S	??	U	??	W	??	Y	(??	[	/??	]	2??	_	9??	a	A??	f	??	o		??	u		??	w
	??	y
	??	{
	&??	}
	0??	
	8?? ?
	B?? ?
	L?? ?		?? ?	?? ?	?? ?	?? ?	?? ?	&?? ?	-?? ?	?? ?	?? ?	?? ?	?? ?	?? ?	&?? ?	0?? ?	?? ?	?? ?	?? ?	?? ?	?? ?	?? ?	?? 	-??		??	??	??	&?? 	??%	?? @		%


  g  nameg  statprof-display?g  documentationf  ?Displays a gprof-like summary of the statistics collected. Unless an
optional @var{port} argument is passed, uses the current output port.? C/R!EQR??M    hH     ]$  :4 5
?$  +4 5
?$  4 54 54 56CCC          g  data
		D g  prior-value		D  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
	??		??			??		??			??	 	??	 	??	"		??	'"	??	(#	??	/$	??	6%	??	>!		?? 		D	   C???        h@   Y  ] 4>  "  G  445 >  "  G  45 6       Q      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?

??		??	'	??	'	??	'	'??	%'	??	2(	??	3(	)??	9(	?? 
		9
  g  nameg  statprof-display-anomolies?g  documentationf  QA sanity check that attempts to detect anomolies in statprof's
statistics.@code{}? C0RD????      h(   6  ] 
?$  4>  "  G  "   ?C  .      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
*
??	,	??	
,	??	-	??	-	??	-	??	%.	?? 		&
  g  nameg  statprof-accumulated-time?g  documentationf  AReturns the time accumulated during the last statprof run.@code{}? CRD??A   h(   0  ] 
?$  4>  "  G  "   C     (      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
0
??	2	??	
2	??	3	??	3	??	3	?? 		#
  g  nameg  statprof-sample-count?g  documentationf  HReturns the number of samples taken during the last statprof run.@code{}? C RMi#RQi$RRi%RSi&RH   h   ?  ] C   z      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
;
?? 		
  g  nameg  statprof-fetch-stacks?g  documentationf  ?Returns a list of stacks, as they were captured since the last call
to @code{statprof-reset}.

Note that stacks are only collected if the @var{full-stacks?} argument
to @code{statprof-reset} is true.? C1Rqs   h8     ] &  C4 5$  45$  4 545?CCC        g  a
		6 g  b		6  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
D	??	
E	??	F	??	G	??	E	??	G	??	"G	??	#H	??	*H	??	1H	?? 		6	  g  nameg  procedure=?? C?R??        h   ?   ] ?4 ?L 5?C       ?       g  tail
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
R	??	S	 ??	T	 ??	T	.??	T	 ??	S	?? 		   Ci??    h   ?   ] ?????C    ?       g  a
		 g  b		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
X	??	X	%??		X	.??	X	"?? 			   C?      h   ?   ]L ?L ??6   ?       g  x
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
[	??	[	 ??	
[	(??	[	?? 		   C? 
   h?   [  ]""  ?(  +4O 5445?45?C?(  ??"???4O 5$  !?4	?????5"??}????? ?"??` 
"??R    S      g  lists
	 ? g  equal?	 ? g  in		 ? g  
n-terminal		 ? g  tails		 ? g  trees			7 g  t		` ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
N
??	O	??	P	??	R	??	R	??	V	??	"V	"??	,V	??	-W	??	6V	??	:Y	??	>P	??	AZ	
??	DZ	??	PZ	??	Q[	??	`P	??	k^	??	n`	??	ua	??	xb	 ??	|b	*??	}b	??	`	?? ?^	?? ?d	
?? ?f	?? ?f	!?? ?f	?? ?f	
?? ?d	?? ?O	?? ?O	,?? ?O	?? #	 ?	  g  nameg  lists->trees? C?R??? h   ?   ] ?C  ?       g  x
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
j	??	j	$?? 		   Cxz~  h    ?   ]44 
556     ?       g  stack
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
h
??	j	
??	m	??	j	
??	i	?? 		  g  nameg  stack->procedures? C?R???H?      h   ?  ] 4455?C   ?      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
o
??	v	??	v	??	v	??	v	?? 		
  g  nameg  statprof-fetch-call-tree?g  documentationf  ?Return a call tree for the previous statprof run.

The return value is a list of nodes, each of which is of the type:
@code
 node ::= (@var{proc} @var{count} . @var{nodes})
@end code? C2R?       h0     ]
 
?$  E4L >   G ? "???       ?       g  i
		) g  result		) g  result			)  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	
?	??	?	??	?	??	?	??	)?	?? 		)	  g  nameg  lp? C    h    ?   ]O  L  Q   L6       ?       g  lp
		  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	?	?? 		
   C???  h@     ] 444L?554L?4L?5??5LL >  "  G  6         g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??		?	&??	?	-??	?	&??	?	??	?	??	?	)??	?	0??	?	0??	#?	7??	%?	0??	&?	-??	'?	&??	)?	??	2?	??	>?	?? 		>
   C/I h(   ?   ] 4>   "  G  4>   "  G   C ?       g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	%?	?? 		'
   C???   h@     ] 444L?554L?4L?5??5LL >  "  G  6         g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??		?	&??	?	-??	?	&??	?	??	?	??	?	)??	?	0??	?	0??	#?	7??	%?	0??	&?	-??	'?	&??	)?	??	2?	??	>?	?? 		>
   C/I h(   ?   ] 4>   "  G  4>   "  G   C ?       g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	%?	?? 		'
   C/I       h(   ?   ] 4>   "  G  4>   "  G   C ?       g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	%?	?? 		'
   C      h?   r  - /   0   3 #  #  	d#  #  O  Q O 4O >   "  G  V4>   X4>   "  G  CX4>   "  G  F     j      g  thunk
	 ? g  loop	 ? g  hz		 ? g  count-calls?		 ? g  full-stacks?		 ? g  thunk		A ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
x
??	A?	?? 	 ?
g  loopS?g  hzS	?g  count-calls?S	?g  full-stacks?S	?   g  nameg  statprof?g  documentationf ?Profile the execution of @var{thunk}, and return its return values.

The stack will be sampled @var{hz} times per second, and the thunk
itself will be called @var{loop} times.

If @var{count-calls?} is true, all procedure calls will be recorded. This
operation is somewhat expensive.

If @var{full-stacks?} is true, at each sample, statprof will store away the
whole call tree, for later analysis. Use @code{statprof-fetch-stacks} or
@code{statprof-fetch-call-tree} to retrieve the last-stored stacks.? CR4Y3[^adg????     hP   M  ](  64?5$   ? &  ??C ?? "???4 5$  CC  E      g  kw
		N g  args		N g  def			N  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	?	??	?	??	?	??	?	??	?	??	?	??	#?	??	&?	
??	-?	??	:?	
??	;?	??	H?	?? 		N	  g  nameg  
kw-arg-ref? C??????    h`     -  1  3 O Q 4 5??4 54 	d54 54 5 
C     
      g  args
			[ g  
kw-arg-ref		[  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	?	??	?	??	%?	??	)?	??	4?	??	5?	
??	A?	??	B?	??	M?	??	N?	??	Z?	?? 			[


   C   h   ?   ]	4 5L 4?6?       g  args
		 g  v			  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?	?
?? 		   Cop    h(   ?  ]	4 5$   O @ 6 ?      g  y
		' g  tmp		'  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
?? 		'  g  documentationf ?Profile the expressions in the body, and return the body's return values.

Keyword arguments:

@table @code
@item #:loop
Execute the body @var{loop} number of times, or @code{#f} for no looping

default: @code{#f}
@item #:hz
Sampling rate

default: @code{20}
@item #:count-calls?
Whether to instrument each function call (expensive)

default: @code{#f}
@item #:full-stacks?
Whether to collect away all sampled stacks into a list

default: @code{#f}
@end table?g  
macro-typeg  defmacro?g  defmacro-argsg  args  C53R???????j@ 
      hx   r  ]  $   C 45 4L 
5  $   "  4455  4>  "  G  ? 	?? 45  	  C    j      g  t
		t g  t
	!	? g  	stop-time
	?	o g  stack	?	o  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?	??	?	??	?	??	?	??	?	??	!?	??	0?	??	4?	??	5?	"??	<?	??	??	??	D?	??	Y ?	??	b?	??	c?	??	i?	??	r?	?? 		t
  g  nameg  gc-callback? CD???@AE{IGHC??F?F?????      h?   ?  ] 
?$  4>  "  G  "   
  
  4 ?5 	L 
 ? ?$  O 45  445 5? 4L >  "  G  445 445 5?>  "  G  CC     ?      g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
	??	?	??	
?	??	?	??	?	??	?	??	#?	??	&?	??	)?	??	,?	??	-?	??	6?	??	:?	??	;?	??	=?	??	@?	??	B?	??	F?	??	J?	??	M?	
??	N?	 ??	T?	
??	U?	#??	Y?	)??	Z?	8??	`?	#??	a?	??	c?	
??	d?	
??	x?	
??	{?	?? ??	,?? ??	<?? ??	,?? ??	(?? ??	
?? %	 ?
   C h8   ?   ]"  # 
?$  C4L >   "  G   ? "???L "???       ?       g  i
		)  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
	??		??				??			??		??	#	??	)	??	)	?? 			1
   CD?F?F???j?@/I       hh   &  ] ? 
?$  =445 5?? 4L >  "  G  	?4
5 ??  "   4>   "  G   C             g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
	??	?	??	?	??	
?	??	?	??	?	??	?	??	?	-??	?	??	?	??	?	??	!?	
??	"?	
??	9 ?	??	; 	??	E 	
??	H	
??	M	??	_	?? 		a
   C       h?   ?  - /   0   3 #  #  O Q O O O Q  Q Q 4>   "  G  V4>   X4>   "  G  CX4>   "  G  F?      g  thunk
	 ? g  loop	 ? g  full-stacks?		 ? g  gc-callback		- ? g  pre		L ? g  thunk		L ? g  post		L ?  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?
?
??	;	?? 	 ?
g  loopS?g  full-stacks?S	?   g  nameg  gcprof?g  documentationf +Do an allocation profile of the execution of @var{thunk}.

The stack will be sampled soon after every garbage collection, yielding
an approximate idea of what is causing allocation in your program.

Since GC does not occur very frequently, you may need to use the
@var{loop} parameter, to cause @var{thunk} to be called @var{loop}
times.

If @var{full-stacks?} is true, at each sample, statprof will store away the
whole call tree, for later analysis. Use @code{statprof-fetch-stacks} or
@code{statprof-fetch-call-tree} to retrieve the last-stored stacks.? C4RC?      g  m
		0  g  filenamef  P/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/statprof.scm?		}
??	4 ?
??	8 ?
??	< ?
??	@ ?
??	D ?
??	H ?
??	L ?
??	P ?
??	T ?
??	U ?	??	X ?
??	\ ?
??? ?
??[ ?
??D ?
?? ?
??? ?
??? ?
??? ?
??? ?
??	? ?
??
? ?
??: ?
??? ?
???
??j
????
???M
???S
??!@m
??%

??(??
??*5?
??.?
??.??
??/??
??0??
??1??
??2l?
??3Q?
??48?
??5??
??>1?
??Al
??B?*
??D[0
??Db6
??Di7
??Dp8
??Dw9
??F;
??GoC
??M>N
??Oh
??P?o
??\?x
??o??
?? ;	o?
   C6 