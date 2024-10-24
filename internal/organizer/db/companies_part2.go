package db

import "github.com/readytotouch/readytotouch/internal/domain"

func companiesPart2() []domain.CompanyProfile {
	return []domain.CompanyProfile{
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Trusting Social",
			URL:  "https://www.trustingsocial.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3506386,
				Alias: "trustingsocial",
				Name:  "Trusting Social",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-TrustingSocial-EI_IE741535.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/TrustingSocial-Reviews-E741535.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4042510486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI fintech company revolutionizing credit scoring using big data technology",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "McAfee",
			URL:  "https://www.mcafee.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2336,
				Alias: "mcafee",
				Name:  "McAfee",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "mcafee",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-McAfee-EI_IE2244.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/McAfee-Reviews-E2244.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4041686516/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Antivirus software company",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hadrian",
			URL:  "https://www.hadrian.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    71668100,
				Alias: "hadrianspace",
				Name:  "Hadrian",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hadrian-EI_IE5364733.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hadrian-Reviews-E5364733.htm",
			},
			OttaProfileSlug:   "Hadrian-2",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/63inUEkF",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Building autonomous factories for aerospace components",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Scaleway",
			URL:  "https://www.scaleway.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9427185,
				Alias: "scaleway",
				Name:  "Scaleway",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "scaleway",
				GoRepositoryCount: 23,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Scaleway-EI_IE1861876.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Scaleway-Reviews-E1861876.htm",
			},
			OttaProfileSlug:   "Scaleway",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/dFlIR1ow",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cloud provider",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Demandbase",
			URL:  "https://www.demandbase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    89759,
				Alias: "demandbase",
				Name:  "Demandbase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Demandbase",
				GoRepositoryCount: 1,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Demandbase-EI_IE271272.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Demandbase-Reviews-E271272.htm",
			},
			OttaProfileSlug:   "Demandbase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/sKrHcCx0",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Go-To-Market engagement & marketing software",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Overwolf",
			URL:  "https://www.overwolf.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    519584,
				Alias: "overwolf.com",
				Name:  "Overwolf",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "overwolf",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Overwolf-EI_IE1963582.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Overwolf-Reviews-E1963582.htm",
			},
			OttaProfileSlug:   "Overwolf",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/duQ9FWwK",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Gaming app & mod development framework",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bitfount",
			URL:  "https://www.bitfount.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69183486,
				Alias: "bitfount",
				Name:  "Bitfount",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bitfount",
				GoRepositoryCount: 0, // Rust 1
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Bitfount",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4044732400/",
					"https://app.welcometothejungle.com/jobs/AGxChIIS",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Federated AI & distributed data science platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Hinge",
			URL:  "https://hinge.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2603651,
				Alias: "hinge-app",
				Name:  "Hinge",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Hinge",
				GoRepositoryCount: 4,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Hinge-EI_IE871901.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Hinge-Reviews-E871901.htm",
			},
			OttaProfileSlug:   "Hinge",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4043430809/",
					"https://app.welcometothejungle.com/jobs/V73cck97",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Dating app for meaningful relationships",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Firework",
			URL:  "https://firework.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18700473,
				Alias: "fireworkhq",
				Name:  "Firework",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Firework-EI_IE2300706.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Firework-Reviews-E2300706.htm",
			},
			OttaProfileSlug:   "Firework",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://app.welcometothejungle.com/jobs/srNX_4cO",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Video-based eCommerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Breakroom",
			URL:  "https://www.breakroom.cc/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18458137,
				Alias: "breakroom",
				Name:  "Breakroom",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "breakroom",
				GoRepositoryCount: 0, // Elixir 13
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Breakroom-EI_IE8002334.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Breakroom-Reviews-E8002334.htm",
			},
			OttaProfileSlug:   "Breakroom",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:    nil,
				domain.Rust:  nil,
				domain.Zig:   nil,
				domain.Scala: nil,
				domain.Elixir: []string{
					"https://www.linkedin.com/jobs/view/4052987616/",
					"https://app.welcometothejungle.com/jobs/p_OpKzJ8",
				},
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "The people-powered job comparison site",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Ansa",
			URL:  "https://www.ansa.dev/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    91158109,
				Alias: "getansa",
				Name:  "Ansa",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Ansa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/EjVpa6gl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Branded customer wallet infrastructure",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "1GLOBAL",
			URL:  "https://www.1global.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    39711,
				Alias: "1global",
				Name:  "1GLOBAL",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-1GLOBAL-EI_IE9805631.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/1GLOBAL-Reviews-E9805631.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048972208/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "eSIM",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Eco-Movement",
			URL:  "https://eco-movement.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15223273,
				Alias: "eco-movement-charge-point-data",
				Name:  "Eco-Movement",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Eco-Movement-EI_IE5353422.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Eco-Movement-Reviews-E5353422.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4046511967/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Charging station data platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Network Optix",
			URL:  "https://www.networkoptix.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1896041,
				Alias:    "network-optix",
				Name:     "Network Optix",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "networkoptix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Network-Optix-EI_IE1103945.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Network-Optix-Reviews-E1103945.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4028418131/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Convert video data into practical insights for businesses and industries of all kinds",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cinemo",
			URL:  "https://www.cinemo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    564661,
				Alias: "cinemo",
				Name:  "Cinemo",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cinemo-EI_IE2189711.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cinemo-Reviews-E2189711.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050001754/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Make every screen an opportunity",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Readdle",
			URL:  "https://readdle.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    629551,
				Alias: "readdle",
				Name:  "Readdle",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "readdle",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Readdle-EI_IE613771.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Readdle-Reviews-E613771.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3971228819/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Creating apps",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Ipsos",
			URL:  "https://www.ipsos.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    4318,
				Alias: "ipsos",
				Name:  "Ipsos",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Ipsos-EI_IE13063.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Ipsos-Reviews-E13063.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4053349624/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Market research company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bedrock Streaming",
			URL:  "https://bedrockstreaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64853943,
				Alias: "bedrock-streaming",
				Name:  "Bedrock Streaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bedrockstreaming",
				GoRepositoryCount: 7,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bedrock-Streaming-EI_IE5955934.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bedrock-Streaming-Reviews-E5955934.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4049210489/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Bedrock creates and operates full-scope streaming platforms",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Whalebone",
			URL:  "https://www.whalebone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10961729,
				Alias: "whaleboneio",
				Name:  "Whalebone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "whalebone",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Whalebone-EI_IE4889910.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Whalebone-Reviews-E4889910.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4007694502/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "OneSignal",
			URL:  "https://onesignal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6424376,
				Alias: "onesignal",
				Name:  "OneSignal",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "OneSignal",
				GoRepositoryCount: 5, // Rust 32
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSignal-EI_IE1952551.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSignal-Reviews-E1952551.htm",
			},
			OttaProfileSlug:   "OneSignal",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/wxzE8h5B",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Customer messaging and engagement platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Metabase",
			URL:  "https://www.metabase.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6460313,
				Alias: "metabase",
				Name:  "Metabase",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "metabase",
				GoRepositoryCount: 0, // Clojure 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Metabase-CA-EI_IE6095700.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Metabase-CA-Reviews-E6095700.htm",
			},
			OttaProfileSlug:   "Metabase",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:     nil,
				domain.Rust:   nil,
				domain.Zig:    nil,
				domain.Scala:  nil,
				domain.Elixir: nil,
				domain.Clojure: []string{
					"https://app.welcometothejungle.com/jobs/K21IaTJ2",
				},
				domain.Haskell: nil,
			},
			ShortDescription: "Open-source Business Intelligence tool",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Anjuna Security",
			URL:  "https://www.anjuna.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18438300,
				Alias: "anjuna-security",
				Name:  "Anjuna Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "anjuna-security",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Anjuna-Security-EI_IE3150191.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Anjuna-Security-Reviews-E3150191.htm",
			},
			OttaProfileSlug:   "Anjuna",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052028718/",
					"https://app.welcometothejungle.com/jobs/yxGdxMhb",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Public-cloud network security",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Aqua Security",
			URL:  "https://www.aquasec.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10034420,
				Alias: "aquasecteam",
				Name:  "Aqua Security",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "aquasecurity",
				GoRepositoryCount: 101,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aqua-Security-Software-EI_IE1785939.11,33.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aqua-Security-Software-Reviews-E1785939.htm",
			},
			OttaProfileSlug:   "Aqua-Security",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4045761623/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud native security platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Verve",
			URL:  "https://verve.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    108605,
				Alias: "verve-ad-solutions",
				Name:  "Verve",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Verve-EI_IE4432646.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Verve-Reviews-E4432646.htm",
			},
			OttaProfileSlug:   "Verve",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4015809270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Real time omnichannel ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Wargaming",
			URL:  "https://wargaming.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    127309,
				Alias: "wargaming-net",
				Name:  "Wargaming",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "wgnet",
				GoRepositoryCount: 3,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wargaming-EI_IE381713.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wargaming-Reviews-E381713.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048943270/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Game developer and publisher",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Napier AI",
			URL:  "https://www.napier.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15197985,
				Alias: "napier",
				Name:  "Napier AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "napier-ai",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Napier-Technologies-EI_IE4370088.11,30.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Napier-Technologies-Reviews-E4370088.htm",
			},
			OttaProfileSlug:   "Napier",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4044948110/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI-driven AML solutions transform financial crime compliance from legal obligation to competitive edge",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Kaseya",
			URL:  "https://www.kaseya.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    21377,
				Alias: "kaseya",
				Name:  "Kaseya",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "kaseya",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Kaseya-EI_IE262966.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Kaseya-Reviews-E262966.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3908613207/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Provider of AI-powered cybersecurity and IT management software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Omio",
			URL:  "https://www.omio.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    2754440,
				Alias: "omio",
				Name:  "Omio",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "omio-labs",
				GoRepositoryCount: 2,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Omio-EI_IE2855962.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Omio-Reviews-E2855962.htm",
			},
			OttaProfileSlug:   "Omio",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4025460474/",
					"https://www.linkedin.com/jobs/view/4058460891/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Comparison & booking website for public transport",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Liebherr",
			URL:  "https://www.liebherr.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11057,
				Alias: "liebherr",
				Name:  "Liebherr Group",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Liebherr-Group-EI_IE787369.11,25.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Liebherr-Group-Reviews-E787369.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3821037210/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Construction equipment manufacturer",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Fiserv",
			URL:  "https://www.fiserv.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3364,
				Alias: "fiserv",
				Name:  "Fiserv",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "fiserv",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fiserv-EI_IE1384.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fiserv-Reviews-E1384.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4026685486/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Gett",
			URL:  "https://www.gett.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1514929,
				Alias: "gettaxi",
				Name:  "Gett",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "GettEngineering",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Gett-EI_IE810731.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Gett-Reviews-E810731.htm",
			},
			OttaProfileSlug:   "Gett",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3819936157/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "SaaS platform for mobility providers",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Stairwell",
			URL:  "https://stairwell.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    68606343,
				Alias: "stairwell-inc",
				Name:  "Stairwell",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "stairwell-inc",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Stairwell",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/hnl2J9ws",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Threat detection & cybersecurity platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Solo.io",
			URL:  "https://www.solo.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11074869,
				Alias: "solo.io",
				Name:  "Solo.io",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "solo-io",
				GoRepositoryCount: 82,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-solo-io-EI_IE5382785.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/solo-io-Reviews-E5382785.htm",
			},
			OttaProfileSlug:   "Solo-io",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/vuvd0QAu",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Delivering API infrastructure software",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Splash Damage",
			URL:  "https://www.splashdamage.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    47756,
				Alias: "splash-damage",
				Name:  "Splash Damage",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "splash-damage",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Splash-Damage-EI_IE470252.11,24.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Splash-Damage-Reviews-E470252.htm",
			},
			OttaProfileSlug:   "SplashDamage",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/A7daqq7I",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Game development studio",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Ditto",
			URL:  "https://ditto.live/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18702497,
				Alias: "dittolive",
				Name:  "Ditto",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "getditto",
				GoRepositoryCount: 6, // Rust 30
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "ditto",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4050904655/",
					"https://app.welcometothejungle.com/jobs/U3QtZINa",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Data sync platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Pinecone",
			URL:  "https://www.pinecone.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    20299330,
				Alias: "pinecone-io",
				Name:  "Pinecone",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "pinecone-io",
				GoRepositoryCount: 6, // Rust 9
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Pinecone-Systems-EI_IE6085804.11,27.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Pinecone-Systems-Reviews-E6085804.htm",
			},
			OttaProfileSlug:   "Pinecone",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/ZudU5aZg",
					"https://www.linkedin.com/jobs/view/4002028518/",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Vector database",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Beyond Identity",
			URL:  "https://www.beyondidentity.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    64665737,
				Alias: "beyond-identity-inc",
				Name:  "Beyond Identity",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gobeyondidentity",
				GoRepositoryCount: 2, // Rust 2
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Beyond-Identity-EI_IE3403008.11,26.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Beyond-Identity-Reviews-E3403008.htm",
			},
			OttaProfileSlug:   "Beyond-Identity",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/akm3fSOi",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Passwordless identity management solutions",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Nightfall AI",
			URL:  "https://www.nightfall.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    19031754,
				Alias: "nightfall-ai",
				Name:  "Nightfall AI",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "nightfallai",
				GoRepositoryCount: 5,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nightfall-AI-EI_IE3097406.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nightfall-AI-Reviews-E3097406.htm",
			},
			OttaProfileSlug:   "Nightfall-AI",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3912841050/",
					"https://app.welcometothejungle.com/jobs/VFnWchKY",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud-native data protection platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Cyberhaven",
			URL:  "https://www.cyberhaven.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10639445,
				Alias: "cyberhaven",
				Name:  "Cyberhaven",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "CyberhavenInc",
				GoRepositoryCount: 13,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cyberhaven-EI_IE2985068.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cyberhaven-Reviews-E2985068.htm",
			},
			OttaProfileSlug:   "Cyberhaven",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/w0d6uKVc",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Data detection & response platform",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Fonoa",
			URL:  "https://www.fonoa.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    30129363,
				Alias: "fonoa",
				Name:  "Fonoa",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "Fonoa-Tech",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Fonoa-EI_IE4125781.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Fonoa-Reviews-E4125781.htm",
			},
			OttaProfileSlug:   "Fonoa",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3960607071/",
					"https://app.welcometothejungle.com/jobs/5Q9cmWeg",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Tax automation for the internet economy",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Payrails",
			URL:  "https://www.payrails.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    75580859,
				Alias: "payrails",
				Name:  "Payrails",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Payrails",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/HCYGHzSQ",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Operating system for payments and financial services",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Teleport",
			URL:  "https://goteleport.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7941233,
				Alias: "go-teleport",
				Name:  "Teleport",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "gravitational",
				GoRepositoryCount: 116,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Teleport-EI_IE1967263.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Teleport-Reviews-E1967263.htm",
			},
			OttaProfileSlug:   "Teleport",
			YouTubeChannelURL: "",
			GoMainLanguage:    true,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4000249330/",
					"https://app.welcometothejungle.com/jobs/vhRdlAui",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Identity-native infrastructure access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Bigblue",
			URL:  "https://www.bigblue.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18776044,
				Alias: "bigblue-co",
				Name:  "Bigblue",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bigblue-EI_IE3379658.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bigblue-Reviews-E3379658.htm",
			},
			OttaProfileSlug:   "Bigblue",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/UU1oL2xl",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Developer of eCommerce logistics tools",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Taxfix",
			URL:  "https://taxfix.de/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10791106,
				Alias: "taxfix",
				Name:  "Taxfix",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "taxfix",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Taxfix-EI_IE1301207.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Taxfix-Reviews-E1301207.htm",
			},
			OttaProfileSlug:   "Taxfix",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://app.welcometothejungle.com/jobs/whJEvleb",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Tax filing app",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Clerk",
			URL:  "https://clerk.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    69336355,
				Alias: "clerkinc",
				Name:  "Clerk.com",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "clerk",
				GoRepositoryCount: 8,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "",
				ReviewsURL:  "",
				Verified:    false,
			},
			OttaProfileSlug:   "Clerk",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Identity tool for React applications",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ClearScore",
			URL:  "https://clearscore.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9320086,
				Alias: "clearscore",
				Name:  "ClearScore",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ClearScore",
				GoRepositoryCount: 2, // Scala 2
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ClearScore-EI_IE1046600.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ClearScore-Reviews-E1046600.htm",
			},
			OttaProfileSlug:   "ClearScore",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://app.welcometothejungle.com/jobs/RmSslgyY",
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Financial wellbeing marketplace",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "OneSchema",
			URL:  "https://www.oneschema.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    74704820,
				Alias: "oneschema",
				Name:  "OneSchema",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "oneschema",
				GoRepositoryCount: 0,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OneSchema-EI_IE7681431.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OneSchema-Reviews-E7681431.htm",
			},
			OttaProfileSlug:   "OneSchema",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://app.welcometothejungle.com/jobs/YvpdZmyz",
				},
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Cloud-based data importing and validation platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "StrongDM",
			URL:  "https://www.strongdm.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9446266,
				Alias: "strongdm",
				Name:  "StrongDM",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "strongdm",
				GoRepositoryCount: 15,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-strongDM-EI_IE4514684.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/strongDM-Reviews-E4514684.htm",
			},
			OttaProfileSlug:   "strongDM",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3978025966/",
					"https://app.welcometothejungle.com/jobs/lys1kK5P",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Manages auditable database access",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "JMA Wireless",
			URL:  "https://jmawireless.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3105986,
				Alias: "jmawireless",
				Name:  "JMA Wireless",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-JMA-Wireless-EI_IE875877.11,23.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/JMA-Wireless-Reviews-E875877.htm",
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048251309/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "JMA Wireless designs and delivers cutting-edge wireless technology",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "ON2IT Cybersecurity",
			URL:  "https://on2it.net/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       489607,
				Alias:    "on2it-cybersecurity",
				Name:     "ON2IT Cybersecurity",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "on2itsecurity",
				GoRepositoryCount: 6,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-ON2IT-EI_IE1075463.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/ON2IT-Reviews-E1075463.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3513136315/", // "We mainly use PHP and Go"
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "CyberSecurity service provider",
			Industries: []domain.Industry{
				domain.IndustryCyberSecurity,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Cycloid",
			URL:  "https://www.cycloid.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10071522,
				Alias: "cycloid",
				Name:  "Cycloid - Sustainable Platform Engineering",
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "cycloidio",
				GoRepositoryCount: 40,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Cycloid-io-EI_IE1562422.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Cycloid-io-Reviews-E1562422.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Cycloid",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4048269201/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Hybrid cloud DevOps platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Circuit",
			URL:  "https://circuit.ai/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       36014538,
				Alias:    "circuitknowledge",
				Name:     "Circuit",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Circuit-TX-EI_IE10004565.11,21.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Circuit-TX-Reviews-E10004565.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3984171089/", // 5+ years of experience with Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "AI-powered platform transforms information into actionable knowledge",
			Industries: []domain.Industry{
				domain.IndustryEdTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Aylo",
			URL:  "https://www.aylo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       98389423,
				Alias:    "ayloservices",
				Name:     "Aylo",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "MindGeekOSS",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Aylo-EI_IE9049055.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Aylo-Reviews-E9049055.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3968183921/", // Familiarity of Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Adult content platforms",
			Industries: []domain.Industry{
				domain.IndustryEntertainment,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "FactSet",
			URL:  "https://www.factset.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       163755,
				Alias:    "factset",
				Name:     "FactSet",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "factset",
				GoRepositoryCount: 1,
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-FactSet-EI_IE6066.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/FactSet-Reviews-E6066.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Factset",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4018413250/", // Proficiency with Go and protobuf for minimum of 1+ years
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Providing instant access to financial data and analytics that investors use to make decisions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Voodoo",
			URL:  "https://voodoo.io/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       5353630,
				Alias:    "voodoo.io",
				Name:     "Voodoo",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "VoodooTeam",
				GoRepositoryCount: 3,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Voodoo-EI_IE1889746.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Voodoo-Reviews-E1889746.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3966903286/", // Our Stack Backend: Go, Node.js
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Company that creates mobile games and apps",
			Industries: []domain.Industry{
				domain.IndustryGameDev,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "home24",
			URL:  "https://www.home24.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2653180,
				Alias:    "home24",
				Name:     "home24",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "home24",
				GoRepositoryCount: 1,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-home24-SE-EI_IE3166011.11,20.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/home24-SE-Reviews-E3166011.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/3872092645/", // Our tech stack includes Golang
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Home & living e-commerce platform",
			Industries: []domain.Industry{
				domain.IndustryECommerce,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Australian Broadcasting Corporation",
			URL:  "https://www.abc.net.au/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2242,
				Alias:    "australian-broadcasting-corporation",
				Name:     "Australian Broadcasting Corporation",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "abcnews",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Australian-Broadcasting-EI_IE416033.11,34.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Australian-Broadcasting-Reviews-E416033.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4050919710/", // Outstanding enterprise-level experience with Golang and extensive background in server-side development
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Independent public broadcaster",
			Industries:       []domain.Industry{
				// Media
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Nitro",
			URL:  "https://www.gonitro.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       76281,
				Alias:    "go-nitro",
				Name:     "Nitro Software, Inc.",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "nitro",
				GoRepositoryCount: 21, // Scala 5
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Nitro-EI_IE402100.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Nitro-Reviews-E402100.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4042171556/", // We are seeking a backend Software Engineer with skills in Scala and Java
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Document productivity SaaS company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Appodeal",
			URL:  "https://appodeal.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       19139082,
				Alias:    "appodeal",
				Name:     "Appodeal",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "appodeal",
				GoRepositoryCount: 2, // Scala 2
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Appodeal-EI_IE4606703.11,19.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Appodeal-Reviews-E4606703.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: nil,
				domain.Rust: []string{
					"https://www.linkedin.com/jobs/view/4052060438/", // Rust Engineer
				},
				domain.Zig: nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4051754807/", // Scala Developer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Tala",
			URL:  "https://tala.co/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2647513,
				Alias:    "talamobile",
				Name:     "Tala",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "inventure",
				GoRepositoryCount: 2, // Scala 2
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Tala-EI_IE1264165.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Tala-Reviews-E1264165.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Tala",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4048262224/", // Senior Scala Developer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Financial platform",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "MacroHealth",
			URL:  "https://macrohealth.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       19012089,
				Alias:    "macrohealth",
				Name:     "MacroHealth",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-MacroHealth-EI_IE3277977.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/MacroHealth-Reviews-E3277977.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4051275869/", // Maintain and extend applications using Scala and Java
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Health Market-as-a-Service platform that enables Payers, Providers, and Health Market Partners",
			Industries:       []domain.Industry{
				// HealthTech, eCommerce?
			},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			// Part of DoorDash
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Wolt",
			URL:  "https://careers.wolt.com/en",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3995846,
				Alias:    "wolt-oy",
				Name:     "Wolt",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "woltapp",
				GoRepositoryCount: 0, // Scala 2
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Wolt-EI_IE1800227.11,15.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Wolt-Reviews-E1800227.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4008937766/", // You have solid experience designing, building, and integrating quality software (microservices) in Scala
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Company making it easy to discover and get restaurants and shops delivered home & to the office",
			Industries:       []domain.Industry{
				// Food delivery
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Bynder",
			URL:  "https://www.bynder.com/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       2427738,
				Alias:    "bynder",
				Name:     "Bynder",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "bynder",
				GoRepositoryCount: 3, // Scala 0
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Bynder-EI_IE1003066.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Bynder-Reviews-E1003066.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Bynder",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4011968287/", // Experience building microservices with containers, Scala and Reactive
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Digital asset management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "OVO",
			URL:  "https://company.ovo.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       643101,
				Alias:    "ovoenergy",
				Name:     "OVO",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "ovotech",
				GoRepositoryCount: 10, // Scala 29
				Verified:          true,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-OVO-EI_IE767881.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/OVO-Reviews-E767881.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4051591843/", // Scala Software Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Clean affordable energy",
			Industries:       []domain.Industry{
				// Energy
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Avetta",
			URL:  "https://www.avetta.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       7584447,
				Alias:    "avetta",
				Name:     "Avetta",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "avetta",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Avetta-EI_IE1182913.11,17.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Avetta-Reviews-E1182913.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4049869381/", // Scala Software Engineer — Backend
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription:          "Contractor risk management platform",
			Industries:                []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Adverty",
			URL:  "https://adverty.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       15233801,
				Alias:    "adverty",
				Name:     "Adverty",
				Verified: false,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Adverty-EI_IE7208426.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Adverty-Reviews-E7208426.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4053026269/", // Backend Scala Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "In-game ad platform",
			Industries: []domain.Industry{
				domain.IndustryAdTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "Unzer",
			URL:  "https://www.unzer.com/en/", // https://www.youpaylater.com/en/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       67571244,
				Alias:    "unzer",
				Name:     "Unzer",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "unzerdev",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Unzer-EI_IE2871480.11,16.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Unzer-Reviews-E2871480.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "Unzer",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4016971905/", // Scala Software Engineer
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "All-in-one payment solutions",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "HERE Technologies",
			URL:  "https://www.here.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       3237134,
				Alias:    "here",
				Name:     "HERE Technologies",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "heremaps",
				GoRepositoryCount: 0, // Scala 2
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-HERE-Technologies-EI_IE753677.11,28.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/HERE-Technologies-Reviews-E753677.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4053090497/", // Software Engineer (Scala/Java)
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Location platform company",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,  // system
			Type: "", // system
			Name: "UBS",
			URL:  "https://www.ubs.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       1214,
				Alias:    "ubs",
				Name:     "UBS",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-UBS-EI_IE3419.11,14.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/UBS-Reviews-E3419.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go:   nil,
				domain.Rust: nil,
				domain.Zig:  nil,
				domain.Scala: []string{
					"https://www.linkedin.com/jobs/view/4034024007/", // Data Engineer — Scala
				},
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "",
			Industries: []domain.Industry{
				domain.IndustryFinTech,
			},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
				domain.Czechia,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "DroneShield",
			URL:  "https://www.droneshield.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       7575557,
				Alias:    "droneshield",
				Name:     "DroneShield",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-DroneShield-EI_IE5175947.11,22.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/DroneShield-Reviews-E5175947.htm",
				Verified:    false,
			},
			OttaProfileSlug:   "",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4032752026/", // Software Engineer — Go
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Counter-UAS solutions",
			Industries:       []domain.Industry{},
			HasEmployeesFromCountries: []domain.Country{
				domain.Ukraine,
			},
		},
		{
			ID:   0,                         // system
			Type: domain.CompanyTypeStartup, // system
			Name: "Automox",
			URL:  "https://www.automox.com/",
			LinkedInProfile: domain.LinkedInProfile{
				ID:       7941257,
				Alias:    "automox",
				Name:     "Automox",
				Verified: true,
			},
			GitHubProfile: domain.GitHubProfile{
				Login:             "automox",
				GoRepositoryCount: 0,
				Verified:          false,
			},
			GlassdoorProfile: domain.GlassdoorProfile{
				OverviewURL: "https://www.glassdoor.com/Overview/Working-at-Automox-EI_IE2236326.11,18.htm",
				ReviewsURL:  "https://www.glassdoor.com/Reviews/Automox-Reviews-E2236326.htm",
				Verified:    true,
			},
			OttaProfileSlug:   "Automox",
			YouTubeChannelURL: "",
			GoMainLanguage:    false,
			Vacancies: domain.Vacancies{
				domain.Go: []string{
					"https://www.linkedin.com/jobs/view/4040310447/",
				},
				domain.Rust:    nil,
				domain.Zig:     nil,
				domain.Scala:   nil,
				domain.Elixir:  nil,
				domain.Clojure: nil,
				domain.Haskell: nil,
			},
			ShortDescription: "Cloud-native automated patch management platform",
			Industries: []domain.Industry{
				domain.IndustryDevOps, // Programming experience in Golang is recommended
			},
			HasEmployeesFromCountries: []domain.Country{},
		},

		// Template
		//{
		//	ID:   0,  // system
		//	Type: "", // system
		//	Name: "",
		//	URL:  "",
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:       0,
		//		Alias:    "",
		//		Name:     "",
		//		Verified: false,
		//	},
		//	GitHubProfile: domain.GitHubProfile{
		//		Login:             "",
		//		GoRepositoryCount: 0,
		//		Verified:          false,
		//	},
		//	GlassdoorProfile: domain.GlassdoorProfile{
		//		OverviewURL: "",
		//		ReviewsURL:  "",
		//		Verified:    false,
		//	},
		//	OttaProfileSlug:   "",
		//	YouTubeChannelURL: "",
		//	GoMainLanguage:    false,
		//	Vacancies: domain.Vacancies{
		//		domain.Go:      nil,
		//		domain.Rust:    nil,
		//		domain.Zig:     nil,
		//		domain.Scala:   nil,
		//		domain.Elixir:  nil,
		//		domain.Clojure: nil,
		//		domain.Haskell: nil,
		//	},
		//	ShortDescription:          "",
		//	Industries:                []domain.Industry{},
		//	HasEmployeesFromCountries: []domain.Country{},
		//},
	}
}