// Code generated by qtc from "organizer.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// organizer-styles.qtpl

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOrganizer(qw422016 *qt422016.Writer, companies []Company, universities []University) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
    <title>Golang companies organizer</title>
    <meta name="description" content="Golang companies organizer">
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="author" type="text/plain" href="https://golang-companies-organizer.readytotouch.com/humans.txt"/>

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`

	<link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&display=swap" rel="stylesheet">

    `)
	streamorganizerStyles(qw422016)
	qw422016.N().S(`
</head>

<body>
<header class="header">
	<div class="header__wrapper">
		<a href="/" class="header__logo">
			<img class="header__logo-img" src="/assets/images/pages/online/logo.svg" alt="logo">
			<h3 class="header__logo-title">Ready To Touch</h3>
		</a>
		<div class="header__stars">
			<iframe src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch&type=star&count=true&size=large" frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
		</div>
	</div>
</header>
<main class="main-wrapper">
    <div class="main-container">
        <section class="organizer">
            <div class="wrapper">
                <div class="organizer__table-container">
                    <div class="table__header-top">
                        <p class="table__result-counter">`)
	qw422016.N().D(len(companies))
	qw422016.N().S(` companies</p>
                    </div>
                    <table class="organizer__table table">
                        <thead class="organizer__head">
                            <tr>
                                <th>
                                    <span>Name</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/common-images/linkedin.svg" alt="linkedin">
                                    <span>LinkedIn</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/online/github-black.svg" alt="github-black">
                                    <span>GitHub</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/common-images/glassdoor.svg" alt="glassdoor">
                                    <span>Glassdoor</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/common-images/similarweb.svg" alt="SimilarWeb">
                                    <span>SimilarWeb</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/common-images/otta.svg" alt="otta">
                                    <span>Otta</span>
                                </th>
                                <th>
                                    <img src="/assets/images/pages/common/link.svg" alt="link">
                                    <span>Vacancies</span>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            `)
	for _, company := range companies {
		qw422016.N().S(`
                            <tr class="table__item">
                                <td>
                                    <div class="table__item name">
                                        <a class="table__item-link" href="`)
		qw422016.E().S(company.URL)
		qw422016.N().S(`">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</a>
                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/square.svg">
                                            <a class="table__item-link" href="https://www.linkedin.com/company/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" title="`)
		qw422016.E().S(company.LinkedInProfile.Name)
		qw422016.N().S(`">Overview</a>
                                        </div>
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/response.svg">
                                            <a class="table__item-link" href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, universities))
		qw422016.N().S(`" title="`)
		qw422016.E().S(company.LinkedInProfile.Name)
		qw422016.N().S(`">Connections</a>
                                        </div>
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/vacancy/briefcase.svg">
                                            <a class="table__item-link" href="`)
		qw422016.E().S(linkedinJobsURL([]Company{company}, golangKeywordsTitles))
		qw422016.N().S(`" title="`)
		qw422016.E().S(company.LinkedInProfile.Name)
		qw422016.N().S(`">Jobs</a>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
`)
		if company.GitHubProfile.Login != "" {
			qw422016.N().S(`                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/square.svg">
                                            <a class="table__item-link" href="https://github.com/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`">Overview</a>
                                        </div>
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/database.svg">
                                            <a class="table__item-link" href="https://github.com/orgs/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`/repositories?q=lang:go">Repositories</a>&nbsp;(`)
			qw422016.N().D(company.GitHubProfile.RepositoryCount)
			qw422016.N().S(`)
                                        </div>
`)
		}
		qw422016.N().S(`                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/square.svg">
                                            <a class="table__item-link" href="`)
		qw422016.E().S(company.GlassdoorProfile.OverviewURL)
		qw422016.N().S(`">Overview</a>
                                        </div>
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/message.svg">
                                            <a class="table__item-link" href="`)
		qw422016.E().S(company.GlassdoorProfile.ReviewsURL)
		qw422016.N().S(`">Reviews</a>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
                                        <div class="table__link-group">
                                            <img src="/assets/images/pages/common/square.svg">
                                            <a class="table__item-link" href="`)
		qw422016.E().S(similarwebURL(company.URL))
		qw422016.N().S(`">Overview</a>
                                        </div>
                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
`)
		if company.OttaProfileSlug != "" {
			qw422016.N().S(`                                            <div class="table__link-group">
                                                <img src="/assets/images/pages/common/square.svg">
                                                <a class="table__item-link" href="https://app.otta.com/companies/`)
			qw422016.E().S(company.OttaProfileSlug)
			qw422016.N().S(`">Overview</a>
                                            </div>
`)
		}
		qw422016.N().S(`                                    </div>
                                </td>
                                <td>
                                    <div class="table__item">
`)
		for i, vacancy := range company.Vacancies {
			qw422016.N().S(`                                            <a class="table__item-link vacancies" href="`)
			qw422016.E().S(vacancy)
			qw422016.N().S(`">Vacancy #`)
			qw422016.N().D(i)
			qw422016.N().S(`</a>
`)
		}
		qw422016.N().S(`                                    </div>
                                </td>
                            </tr>
                            `)
	}
	qw422016.N().S(`
                        </tbody>
                    </table>
                </div>
                <div class="organizer__linkedin">
                    <h2 class="headline headline--lvl1 organizer__block-title">LinkedIn</h2>
                    <div class="organizer__links">
                        <a class="organizer__link" href="`)
	qw422016.E().S(linkedinConnectionsURL(companies, nil))
	qw422016.N().S(`">LinkedIn Connections [Companies]</a>
                        `)
	if len(universities) > 0 {
		qw422016.N().S(`
                            <a class="organizer__link" href="`)
		qw422016.E().S(linkedinConnectionsURL(companies, universities))
		qw422016.N().S(`">LinkedIn Connections [Companies] [Universities]</a>
                        `)
	}
	qw422016.N().S(`
                        <a class="organizer__link" href="`)
	qw422016.E().S(linkedinJobsURL(companies, golangKeywordsTitles))
	qw422016.N().S(`">LinkedIn Jobs [Companies] [Worldwide]</a>
                        <a class="organizer__link" href="`)
	qw422016.E().S(linkedinJobsURL(nil, golangKeywordsTitles))
	qw422016.N().S(`">LinkedIn Jobs [Worldwide]</a>
                    </div>
                </div>
            </div>
        </section>
    </div>
</main>
<footer class="footer">
    <div class="footer__wrapper main-container">
        <div class="footer__group-u8">
            <div class="footer__info">
                <a href="/" class="footer__title">
                    <h3 class="footer__title-link">Ready To Touch</h3>
                </a>
                <p class="footer__subtitle">Anonymous job search</p>
                <div class="footer__map">
                    <img class="footer__map-Ukraine" src="/assets/images/pages/online/map-of-Ukraine.png" alt="Map of Ukraine">
                </div>
            </div>
            <div class="footer__middle-section">
                <iframe class="footer__stars" src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch&type=star&count=true&size=large" frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                <a href="https://github.com/readytotouch/golang-companies-organizer" target="_blank" class="footer__statistics-link">View the project on GitHub</a>
                <a href="https://readytotouch.com/" target="_blank" class="footer__github-link">readytotouch.com</a>
            </div>
            <div class="footer__support">
                <a href="https://savelife.in.ua/en/" target="_blank" class="footer__link">
                    <figure class="footer__figure">
                        <img width="60" height="24" src="https://savelife.in.ua/wp-content/themes/savelife/assets/images/new-logo-en.svg" alt="support">
                        <figcaption class="footer__caption">Support</figcaption>
                        <img src="/assets/images/pages/online/arrow-up.svg" alt="arrow">
                    </figure>
                </a>
                <a href="https://war.ukraine.ua/" target="_blank" class="footer__link">
                    <figure class="footer__figure">
                        <figcaption class="footer__caption">war.ukraine.ua</figcaption>
                        <img src="/assets/images/pages/online/arrow-up.svg"  alt="arrow">
                    </figure>
                </a>
            </div>
        </div>
        <div class="footer__copyrights">
            <span>© 2024 Yaroslav Podorvanov</span>
            <img class="footer__flag-UA" src="/assets/images/pages/online/flag-of-Ukraine.svg" alt="Flag of Ukraine">
        </div>
    </div>
</footer>

</body>
</html>
`)
}

func WriteOrganizer(qq422016 qtio422016.Writer, companies []Company, universities []University) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizer(qw422016, companies, universities)
	qt422016.ReleaseWriter(qw422016)
}

func Organizer(companies []Company, universities []University) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizer(qb422016, companies, universities)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}