package utils

var stationList = map[string]string{
	"abrantes":                            "94-52001",
	"ademia":                              "94-36046",
	"afife":                               "94-18119",
	"agualva - cacem":                     "94-61002",
	"aguas santas - palmilheira":          "94-3046",
	"aguda":                               "94-39057",
	"agueda":                              "94-42218",
	"aguieira":                            "94-42267",
	"aguim":                               "94-37093",
	"albergaria dos doze":                 "94-34439",
	"albufeira - ferreiras":               "94-78063",
	"alcacovas":                           "94-74120",
	"alcaide":                             "94-53504",
	"alcains":                             "94-53140",
	"alcantara - mar":                     "94-69039",
	"alcantara - terra":                   "94-67025",
	"alcantarilha":                        "94-90092",
	"alcaria":                             "94-53629",
	"aldeia":                              "94-49411",
	"alegria":                             "94-11064",
	"alferrarede":                         "94-52068",
	"alges":                               "94-69088",
	"algoz":                               "94-90050",
	"algueirao - mem martins":             "94-61069",
	"alhandra":                            "94-31237",
	"alhos vedros":                        "94-95075",
	"almancil":                            "94-78253",
	"almourol":                            "94-51102",
	"alpedrinha":                          "94-53355",
	"alvaraes":                            "94-6338",
	"alvega - ortiga":                     "94-52209",
	"alverca":                             "94-31187",
	"alvito":                              "94-74351",
	"amadora":                             "94-60087",
	"ameal":                               "94-35097",
	"amoreiras - odemira":                 "94-77099",
	"ancora-praia":                        "94-18150",
	"aregos":                              "94-9191",
	"areia - darque":                      "94-6395",
	"arentim":                             "94-29066",
	"areosa":                              "94-18044",
	"arrifana":                            "94-44248",
	"arronches":                           "94-57174",
	"arroteia":                            "94-21055",
	"assumar":                             "94-57117",
	"avanca":                              "94-38216",
	"aveiro":                              "94-38000",
	"aveleda":                             "94-29108",
	"azambuja":                            "94-33001",
	"azurva":                              "94-42051",
	"badajoz":                             "71-37606",
	"baixa da banheira":                   "94-95059",
	"baracal":                             "94-48454",
	"barca da amieira - envendos":         "94-52415",
	"barcelos":                            "94-6122",
	"barqueiros":                          "94-9324",
	"barquinha":                           "94-51045",
	"barragem de belver":                  "94-52241",
	"barreiro":                            "94-95000",
	"barreiro-a":                          "94-95026",
	"barrimau":                            "94-5058",
	"barroselas":                          "94-6304",
	"beja":                                "94-75002",
	"belem":                               "94-69054",
	"belmonte - manteigas":                "94-54197",
	"belver":                              "94-52282",
	"bemposta":                            "94-55129",
	"bencanta":                            "94-35170",
	"benespera":                           "94-54338",
	"benfica":                             "94-60046",
	"benquerencas":                        "94-52878",
	"bifurcacao de lares":                 "94-64022",
	"bobadela":                            "94-31070",
	"boliqueime":                          "94-78147",
	"bom joao":                            "94-73031",
	"bombarral":                           "94-62703",
	"braco de prata":                      "94-31005",
	"braga":                               "94-29157",
	"bustelo":                             "94-8334",
	"cabanoes":                            "94-42150",
	"cabeda":                              "94-8029",
	"cacela":                              "94-73452",
	"cacia":                               "94-38075",
	"cais do sodre":                       "94-69005",
	"caldas da rainha":                    "94-63008",
	"caldas de moledo":                    "94-9399",
	"caminha":                             "94-18242",
	"campolide":                           "94-60004",
	"canas - felgueira":                   "94-46599",
	"canelas":                             "94-38117",
	"canicos":                             "94-28100",
	"carapecos":                           "94-6197",
	"carcavelos":                          "94-69187",
	"caria":                               "94-54148",
	"carrascal - delongo":                 "94-40030",
	"carreco":                             "94-18085",
	"carregado":                           "94-31336",
	"carregal do sal":                     "94-46482",
	"carreira":                            "94-6056",
	"carvalha":                            "94-18416",
	"carvalhal portela":                   "94-42309",
	"carvalheira - maceda":                "94-38356",
	"carvalhos de figueiredo":             "94-40113",
	"casa branca":                         "94-74005",
	"casais":                              "94-35147",
	"casal do alvaro":                     "94-42168",
	"cascais":                             "94-69260",
	"castanheira do ribatejo":             "94-31310",
	"castelejo":                           "94-46409",
	"castelo branco":                      "94-53009",
	"castelo novo":                        "94-53314",
	"castro marim":                        "94-73502",
	"cavaco":                              "94-44172",
	"caxarias":                            "94-34330",
	"caxias":                              "94-69120",
	"caíde":                               "94-8383",
	"celorico da beira":                   "94-48405",
	"cerdeira":                            "94-49205",
	"cete":                                "94-8227",
	"chanca":                              "94-56101",
	"chao de macas - fatima":              "94-34249",
	"coimbra-b":                           "94-36004",
	"coimbroes":                           "94-39149",
	"conceicao":                           "94-73379",
	"contumil":                            "94-3004",
	"cortegaca":                           "94-38372",
	"couto de cambeses":                   "94-29033",
	"couto de cucujaes":                   "94-44297",
	"covas":                               "94-28290",
	"covelinhas":                          "94-10090",
	"covilha":                             "94-54007",
	"crato":                               "94-56267",
	"cruz quebrada":                       "94-69104",
	"cuba":                                "94-74476",
	"cuca":                                "94-28209",
	"curia":                               "94-37119",
	"curvaceiras":                         "94-40063",
	"dagorda - peniche":                   "94-62802",
	"darque":                              "94-6387",
	"dois portos":                         "94-62380",
	"donas":                               "94-53520",
	"durraes":                             "94-6262",
	"eirol":                               "94-42119",
	"eixo":                                "94-42077",
	"elvas":                               "94-57497",
	"entrecampos":                         "94-66050",
	"entroncamento":                       "94-34009",
	"ermesinde":                           "94-4002",
	"ermida":                              "94-9258",
	"ermidas - sado":                      "94-93005",
	"escapaes":                            "94-44222",
	"esgueira":                            "94-42028",
	"esmeriz":                             "94-5033",
	"esmoriz":                             "94-38406",
	"espadanal da azambuja":               "94-31401",
	"espadaneira":                         "94-35162",
	"espinho":                             "94-39008",
	"espinho - vouga":                     "94-44016",
	"esqueiro":                            "94-18291",
	"estarreja":                           "94-38158",
	"estombar - lagoa":                    "94-90241",
	"estoril":                             "94-69245",
	"evora":                               "94-83006",
	"famalicao":                           "94-5074",
	"famalicao da nazare":                 "94-63180",
	"faria":                               "94-44271",
	"faro":                                "94-73007",
	"fatela - penamacor":                  "94-53462",
	"feliteira":                           "94-62364",
	"fernando po":                         "94-71050",
	"ferradosa":                           "94-11114",
	"ferragudo":                           "94-90274",
	"ferrao":                              "94-10165",
	"ferreiros":                           "94-29132",
	"figueira da foz":                     "94-64113",
	"fontela":                             "94-64089",
	"fontela-a":                           "94-64097",
	"formoselha - santo varao":            "94-35030",
	"fornos de algodres":                  "94-48249",
	"francelos":                           "94-39081",
	"fratel":                              "94-52571",
	"freineda":                            "94-49387",
	"freixo de numao - mos do douro":      "94-11247",
	"funcheira":                           "94-77008",
	"fundao":                              "94-53546",
	"fungalvaz":                           "94-34199",
	"fuseta":                              "94-73205",
	"fuseta-a":                            "94-73197",
	"gata":                                "94-49049",
	"general torres":                      "94-39172",
	"giesteira":                           "94-28159",
	"godim":                               "94-9423",
	"gondarem":                            "94-18325",
	"gouveia":                             "94-48165",
	"grandola":                            "94-92577",
	"granja":                              "94-39040",
	"granja do ulmeiro - alfarelos":       "94-35006",
	"guarda":                              "94-49007",
	"guia":                                "94-63800",
	"guimaraes":                           "94-24000",
	"hospital sao joao":                   "94-21030",
	"irivo":                               "94-8243",
	"jerumelo":                            "94-62257",
	"juncal":                              "94-9050",
	"lagos":                               "94-90464",
	"lamarosa":                            "94-34090",
	"lapa":                                "94-44057",
	"lapa do lobo":                        "94-46581",
	"lardosa":                             "94-53215",
	"lares":                               "94-64055",
	"lavradio":                            "94-95042",
	"leandro":                             "94-4036",
	"leca do balio":                       "94-21071",
	"leiria":                              "94-63560",
	"lisboa oriente":                      "94-31039",
	"lisboa rossio":                       "94-59006",
	"lisboa santa apolonia":               "94-30007",
	"litem":                               "94-34504",
	"livracao":                            "94-8474",
	"livramento":                          "94-73239",
	"lordelo":                             "94-28191",
	"loule":                               "94-78238",
	"lourical":                            "94-63875",
	"louro":                               "94-5116",
	"lousado":                             "94-5009",
	"luso - bucaco":                       "94-46094",
	"luz":                                 "94-73262",
	"macainhas":                           "94-54262",
	"macinhata":                           "94-42325",
	"madalena":                            "94-39123",
	"mafra":                               "94-62166",
	"malveira":                            "94-62224",
	"mangualde":                           "94-48009",
	"marco de canaveses":                  "94-9001",
	"marinha grande":                      "94-63461",
	"martinganca - maceira":               "94-63404",
	"marujal":                             "94-65052",
	"marvila":                             "94-66019",
	"massama - barcarena":                 "94-60137",
	"mato de miranda":                     "94-32383",
	"mazagao":                             "94-29116",
	"mealhada":                            "94-37051",
	"meia praia":                          "94-90431",
	"meinedo":                             "94-8359",
	"merces":                              "94-61051",
	"messines - alte":                     "94-77735",
	"mexilhoeira grande":                  "94-90381",
	"midoes":                              "94-6080",
	"mira sintra - melecas":               "94-62042",
	"miramar":                             "94-39073",
	"mirao":                               "94-9225",
	"miuzela":                             "94-49239",
	"mogofores":                           "94-37143",
	"moimenta - alcafache":                "94-46763",
	"moita":                               "94-95109",
	"moledo do minho":                     "94-18200",
	"monte abraao":                        "94-60111",
	"monte de lobos":                      "94-46219",
	"monte de paramos":                    "94-44040",
	"monte estoril":                       "94-69252",
	"monte gordo":                         "94-73544",
	"monte real":                          "94-63685",
	"monte redondo":                       "94-63735",
	"montemor":                            "94-65029",
	"mortagua":                            "94-46243",
	"moscavide":                           "94-31047",
	"mosteiro":                            "94-9134",
	"mouquim":                             "94-5108",
	"mourisca do vouga":                   "94-42259",
	"mouriscas - a":                       "94-52167",
	"nelas":                               "94-46672",
	"nespereira":                          "94-28266",
	"nine":                                "94-6007",
	"obidos":                              "94-62836",
	"oeiras":                              "94-69179",
	"oia":                                 "94-37275",
	"oleiros":                             "94-8250",
	"olhao":                               "94-73106",
	"oliveira":                            "94-8409",
	"oliveira de azemeis":                 "94-44339",
	"oliveira do bairro":                  "94-37218",
	"oliveirinha - cabanas":               "94-46524",
	"oronhe":                              "94-42176",
	"outeiro":                             "94-62612",
	"ovar":                                "94-38299",
	"paco de arcos":                       "94-69146",
	"pacos de brandao":                    "94-44107",
	"paialvo":                             "94-34157",
	"pala":                                "94-9100",
	"palmela":                             "94-68080",
	"pampilhosa":                          "94-37002",
	"papízios":                            "94-46441",
	"parada":                              "94-8201",
	"paraimo - sangalhos":                 "94-37184",
	"paramos":                             "94-38414",
	"parede":                              "94-69203",
	"paredes":                             "94-8276",
	"parque das cidades":                  "94-78295",
	"pataias":                             "94-63354",
	"paul":                                "94-62745",
	"pedra furada":                        "94-62133",
	"pegoes":                              "94-71126",
	"pelariga":                            "94-34702",
	"penafiel":                            "94-8318",
	"penteado":                            "94-95125",
	"pereira":                             "94-35055",
	"pereirinhas":                         "94-28217",
	"pero negro":                          "94-62315",
	"pinhal novo":                         "94-68007",
	"pinhao":                              "94-10249",
	"poceirao":                            "94-71001",
	"pocinho":                             "94-12005",
	"poco barreto":                        "94-90134",
	"pombal":                              "94-34645",
	"ponte de sor":                        "94-55293",
	"porta nova":                          "94-73338",
	"portalegre":                          "94-57000",
	"portela":                             "94-4101",
	"portela de sintra":                   "94-61093",
	"portimao":                            "94-90290",
	"porto campanha":                      "94-2006",
	"porto rei":                           "94-9282",
	"porto sao bento":                     "94-1008",
	"povoa":                               "94-31146",
	"praca do quebedo":                    "94-68130",
	"pragal":                              "94-17087",
	"praia do ribatejo":                   "94-51128",
	"praias do sado-a":                    "94-91058",
	"queluz - belas":                      "94-60103",
	"quintans":                            "94-37358",
	"ramalhal":                            "94-62547",
	"reboleira":                           "94-60095",
	"recarei - sobreira":                  "94-8177",
	"recesinhos":                          "94-8441",
	"rede":                                "94-9357",
	"regua":                               "94-10009",
	"reguengo - vale da pedra - pontevel": "94-33084",
	"retaxo":                              "94-52845",
	"reveles":                             "94-65128",
	"riachos - torres novas - golega":     "94-32466",
	"rio de mouro":                        "94-61044",
	"rio meao":                            "94-44115",
	"rio tinto":                           "94-3038",
	"rochoso":                             "94-49163",
	"roma - areeiro":                      "94-66035",
	"runa":                                "94-62422",
	"ruílhe":                              "94-29074",
	"sabugal":                             "94-54429",
	"sabugo":                              "94-62091",
	"sacavem":                             "94-31062",
	"salir do porto":                      "94-63107",
	"salreu":                              "94-38125",
	"sampaio - oleiros":                   "94-44073",
	"sanfins":                             "94-44198",
	"santa cita":                          "94-40105",
	"santa clara - saboia":                "94-77388",
	"santa comba dao":                     "94-46367",
	"santa cruz - damaia":                 "94-60038",
	"santa cruz-damaia":                   "94-60038",
	"santa eulalia":                       "94-57315",
	"santa iria":                          "94-31112",
	"santa margarida":                     "94-51185",
	"santana - cartaxo":                   "94-32045",
	"santarem":                            "94-32185",
	"santiago de riba - ul":               "94-44313",
	"santo amaro":                         "94-69161",
	"santo tirso":                         "94-28068",
	"santos":                              "94-69013",
	"sao frutuoso":                        "94-4051",
	"sao gemil":                           "94-21006",
	"sao joao da madeira":                 "94-44255",
	"sao joao das craveiras":              "94-71159",
	"sao joao de ver":                     "94-44156",
	"sao joao do estoril":                 "94-69237",
	"sao joao loure":                      "94-42093",
	"sao mamede":                          "94-62786",
	"sao mamede de infesta":               "94-21048",
	"sao martinho do campo":               "94-8102",
	"sao martinho do porto":               "94-63131",
	"sao pedro da torre":                  "94-18440",
	"sao pedro do estoril":                "94-69229",
	"sao romao":                           "94-4077",
	"sapataria":                           "94-62299",
	"sarnadas":                            "94-52803",
	"seica - ourem":                       "94-34264",
	"seixas":                              "94-18275",
	"senhora da agonia":                   "94-18234",
	"senhora das neves":                   "94-6312",
	"sernada do vouga":                    "94-43000",
	"sete rios":                           "94-66076",
	"setil":                               "94-32003",
	"setubal":                             "94-68122",
	"silva":                               "94-6155",
	"silvalde":                            "94-38422",
	"silvalde - vouga":                    "94-44032",
	"silves":                              "94-90183",
	"simoes":                              "94-34744",
	"sintra":                              "94-61101",
	"soalheira":                           "94-53264",
	"soito":                               "94-46185",
	"soudos - vila nova":                  "94-40014",
	"soure":                               "94-34801",
	"souselas":                            "94-36087",
	"suzao":                               "94-8060",
	"tadim":                               "94-29090",
	"taipa - requeixo":                    "94-42127",
	"tamel":                               "94-6213",
	"tancos":                              "94-51086",
	"taveiro":                             "94-35139",
	"tavira":                              "94-73320",
	"terronhas":                           "94-8136",
	"tojeirinha":                          "94-52738",
	"tomar":                               "94-40154",
	"torre das vargens":                   "94-56002",
	"torres vedras":                       "94-62471",
	"tortosendo":                          "94-53678",
	"tramagal":                            "94-51243",
	"trancoso":                            "94-8151",
	"travagem":                            "94-4010",
	"travasso":                            "94-42135",
	"trofa":                               "94-4630",
	"tua":                                 "94-11007",
	"tunes":                               "94-78006",
	"vacarica":                            "94-46045",
	"valadares":                           "94-39115",
	"valado - nazare - alcobaca":          "94-63263",
	"vale de figueira":                    "94-32284",
	"vale de prazeres":                    "94-53397",
	"vale de santarem":                    "94-32102",
	"valega":                              "94-38240",
	"valenca":                             "94-7005",
	"valongo":                             "94-8086",
	"valongo - vouga":                     "94-42291",
	"vargelas":                            "94-11148",
	"venda do alcaide":                    "94-68049",
	"vendas novas":                        "94-72009",
	"vermoil":                             "94-34553",
	"verride":                             "94-65086",
	"vesuvio":                             "94-11197",
	"viana do castelo":                    "94-18002",
	"vigo-guixar":                         "71-22308",
	"vila da feira":                       "94-44206",
	"vila das aves":                       "94-28134",
	"vila fernando":                       "94-49114",
	"vila franca das naves":               "94-48546",
	"vila franca de xira":                 "94-31278",
	"vila mea":                            "94-8433",
	"vila nova da baronia":                "94-74278",
	"vila nova da rainha":                 "94-31377",
	"vila nova de ancos":                  "94-34868",
	"vila nova de cerveira":               "94-18341",
	"vila nova de gaia - devesas":         "94-39164",
	"vila pouca do campo":                 "94-35105",
	"vila real de santo antonio":          "94-73569",
	"vila velha de rodao":                 "94-52647",
	"vilar formoso":                       "94-49460",
	"vilela - fornos":                     "94-36053",
	"virtudes":                            "94-33043",
	"vizela":                              "94-28233",
	"zibreira":                            "94-62331",
}

type TrainStationsInfo []struct {
	Delay       int `json:"delay"`
	TrainOrigin struct {
		Code        string `json:"code"`
		Designation string `json:"designation"`
	} `json:"trainOrigin"`
	TrainDestination struct {
		Code        string `json:"code"`
		Designation string `json:"designation"`
	} `json:"trainDestination"`
	DepartureTime string `json:"departureTime"`
	ArrivalTime   string `json:"arrivalTime"`
	TrainNumber   int    `json:"trainNumber"`
	TrainService  struct {
		Code        string `json:"code"`
		Designation string `json:"designation"`
	} `json:"trainService"`
	Platform  string      `json:"platform"`
	Occupancy interface{} `json:"occupancy"`
	Eta       string      `json:"eta"`
	Etd       string      `json:"etd"`
}

type TrainInfo struct {
	TrainNumber int `json:"trainNumber"`
	ServiceCode struct {
		Code        string `json:"code"`
		Designation string `json:"designation"`
	} `json:"serviceCode"`
	Delay      int         `json:"delay"`
	Occupancy  interface{} `json:"occupancy"`
	Latitude   string      `json:"latitude"`
	Longitude  string      `json:"longitude"`
	Status     string      `json:"status"`
	TrainStops []struct {
		Station struct {
			Code        string `json:"code"`
			Designation string `json:"designation"`
		} `json:"station"`
		Arrival   interface{} `json:"arrival"`
		Departure string      `json:"departure"`
		Platform  string      `json:"platform"`
		Latitude  string      `json:"latitude"`
		Longitude string      `json:"longitude"`
		Delay     int         `json:"delay"`
		Eta       interface{} `json:"eta"`
		Etd       string      `json:"etd"`
	} `json:"trainStops"`
}

type LiveStatus struct {
	Trains     []TrainInfo `json:"trains"`
	Updated_at int64       `json:"updated_at"`
}
