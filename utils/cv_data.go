package utils

import "cv-website/model"

var CVData = model.PageData{
	Profile: model.Profile{
		Name:     "Mohammed Najib Benghuzzi",
		Title:    "Self-taught Developer · RHCSA · Flutter - Golang Developer · IT Engineer",
		Company:  "Almadar Aljadid",
		Location: "Benghazi, Libya",
		Bio: `Self-taught developer and IT engineer with a passion for building clean, maintainable software. RHCSA certified with hands-on experience in Linux system administration, containerization, and full-stack development. Working in IT Technical Support at Almadar Aljadid has not stopped me from continuously learning and building with modern technologies from Flutter and Golang to Docker, Kubernetes, and beyond.`,
		Photo: "assets/images/mnb.jpg",
		Links: []model.Link{
			{Label: "Email", URL: "mailto:mnbenghuzzi@gmail.com"},
			{Label: "Phone", URL: "tel:0925566287"},
			{Label: "SMS", URL: "sms:0925566287"},
		},
	},

	Experience: []model.Experience{
		{
			Company: "Almadar Aljadid",
			Role:    "IT Technical Support",
			Start:   "Present",
			End:     "",
			Highlights: []string{
				"Provide IT infrastructure and technical support across the organization",
				"Manage and troubleshoot Active Directory, MDT deployments, Linux-based systems, and network services",
				"Support end-users and maintain IT operations for critical business functions",
			},
		},
	},

	Skills: []model.SkillGroup{
		{
			Category: "Infrastructure & DevOps",
			Items:    []string{"Linux", "Docker", "Kubernetes", "Virtualization"},
		},
		{
			Category: "Languages & Frameworks",
			Items:    []string{"Flutter", "Golang", "PowerShell", "Frontend Development", "Backend Development"},
		},
		{
			Category: "Databases & Data",
			Items:    []string{"PostgreSQL", "Firebase", "DBML", "SQL"},
		},
		{
			Category: "Practices & Tools",
			Items:    []string{"Clean Architecture", "TDD", "CI/CD", "Git", "GitHub", "Cloud Storage"},
		},
		{
			Category: "Design",
			Items:    []string{"Figma", "Adobe XD"},
		},
		{
			Category: "Traits",
			Items:    []string{"Problem Solver", "Quick Learner", "Embraces Challenges", "IoT"},
		},
	},

	Education: []model.Education{
		{
			Institution: "Red Hat, Inc.",
			Degree:      "Red Hat Certified System Administrator (RHCSA)",
			Year:        "2019",
			Image:       "assets/images/redhat_cert.png",
			VerifyURL:   "https://rhtapps.redhat.com/verify?certId=190-192-609",
		},
	},
}
