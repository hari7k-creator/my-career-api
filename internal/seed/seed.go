package seed

import (
	"github.com/hari7k-creator/my-career-api/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeedData creates initial sample careers for the database
func SeedData(client *mongo.Client) error {
	log.Println("Seeding database with sample careers...")

	collection := client.Database("career_arch").Collection("careers")

	// Check if data already exists
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("Careers already exist in database, skipping seed")
		return nil
	}

	// Create sample careers
	careers := []models.Career{
		createUPSKCareer(),
		createAICareer(),
		createFinanceCareer(),
		createDesignCareer(),
	}

	// Insert careers
	for _, career := range careers {
		career.CreatedAt = time.Now().Format(time.RFC3339)
		career.UpdatedAt = career.CreatedAt

		_, err := collection.InsertOne(context.Background(), career)
		if err != nil {
			return err
		}
	}

	log.Printf("Seeded %d careers successfully", len(careers))
	return nil
}

func createUPSKCareer() models.Career {
	return models.Career{
		Name:        "Indian Administrative Service (IAS)",
		Slug:        "ias",
		Description: "Leadership at district, state, and central government tiers. Manage law enforcement policy, public budget allocation, and macro governance implementation.",
		Category:    "civil",
		Difficulty:  "elite",
		Tags:        []string{"government", "administration", "policy", "leadership"},
		Skills:      []string{"Public Administration", "Constitutional Law", "Crisis Governance", "Policy Analysis", "Stakeholder Management"},
		ContentBlocks: []models.ContentBlock{
			{
				Type:   "hero",
				Title:  "Indian Administrative Service",
				Content: map[string]interface{}{
					"subtitle":   "The premier civil service of India",
					"description": "IAS officers form the backbone of India's administrative machinery, serving at block, district, state, and central levels. They play a crucial role in policy formulation and implementation.",
				},
			},
			{
				Type:   "requirements",
				Title:  "Eligibility & Requirements",
				Content: map[string]interface{}{
					"requirements": []map[string]interface{}{
						{
							"title":       "Citizenship",
							"description": "Indian citizen (NCN/GCN applicable for certain posts)",
							"icon":        "🇮🇳",
						},
						{
							"title":       "Age Limit",
							"description": "21-32 years (relaxation for reserved categories)",
							"icon":        "📅",
						},
						{
							"title":       "Educational Qualification",
							"description": "Bachelor's degree from any recognized university",
							"icon":        "🎓",
						},
					},
				},
			},
			{
				Type:   "assessment",
				Title:  "UPSC Civil Services Examination",
				Content: map[string]interface{}{
					"stages": []map[string]interface{}{
						{
							"name":        "Preliminary Examination",
							"description": "Objective type screening test (2 papers)",
							"marks":       "400 (qualifying)",
							"subjects":    []string{"General Studies Paper I", "CSAT Paper II"},
						},
						{
							"name":        "Main Examination",
							"description": "Descriptive type written examination (9 papers)",
							"marks":       "1750",
							"subjects":    []string{"Essay", "GS Papers I-IV", "Optional Papers I-II", "Language Papers"},
						},
						{
							"name":        "Personality Test",
							"description": "Interview/Personality Assessment",
							"marks":       "275",
							"subjects":    []string{"Interview Assessment"},
						},
					},
				},
			},
			{
				Type:   "roadmap",
				Title:  "Preparation Roadmap",
				Content: map[string]interface{}{
					"timeEstimate": "12-18 months",
					"steps": []models.RoadmapStep{
						{
							Title:           "Understand Exam Pattern & Syllabus",
							Description:     "Thoroughly study the UPSC CSE syllabus and previous year question papers",
							EstimatedMonths: 1,
							Status:          "completed",
							Resources: []models.Resource{
								{Title: "UPSC CSE Official Notification", URL: "https://upsc.gov.in", Type: "documentation"},
							},
						},
						{
							Title:           "Build Foundation with NCERTs",
							Description:     "Read NCERT books from Class 6-12 for basic understanding",
							EstimatedMonths: 3,
							Status:          "in-progress",
						},
						{
							Title:           "Standard Reference Books",
							Description:     "Study standard reference books for each subject",
							EstimatedMonths: 4,
							Status:          "todo",
						},
						{
							Title:           "Current Affairs",
							Description:     "Regularly read newspapers and prepare notes",
							EstimatedMonths: 6,
							Status:          "todo",
						},
						{
							Title:           "Answer Writing Practice",
							Description:     "Start practicing answer writing for mains",
							EstimatedMonths: 8,
							Status:          "todo",
						},
						{
							Title:           "Mock Tests & Revision",
							Description:     "Take mock tests and revise all topics",
							EstimatedMonths: 10,
							Status:          "todo",
						},
					},
				},
			},
			{
				Type:   "metrics",
				Title:  "Career Metrics",
				Content: map[string]interface{}{
					"metrics": []map[string]interface{}{
						{"label": "Selection Ratio", "value": "1:1000", "description": "Approximately 1000 candidates clear prelims for each vacancy"},
						{"label": "Duration", "value": "12-18 months", "description": "Typical preparation time required"},
						{"label": "Starting Salary", "value": "₹56,100", "description": "Pay Level 10 as per 7th CPC"},
						{"label": "Growth", "value": "Cabinet Secretary", "description": "Highest post achievable"},
					},
				},
			},
		},
		Roadmap: []models.RoadmapStep{
			{Title: "Prelims Preparation", Description: "Complete Prelims syllabus and practice mocks", EstimatedMonths: 6},
			{Title: "Mains Preparation", Description: "Complete Mains syllabus and answer writing", EstimatedMonths: 6},
			{Title: "Interview Prep", Description: "Practice GDs and personal interviews", EstimatedMonths: 3},
		},
		Metrics: models.Metrics{
			DurationMonths: 12,
			EntryThreshold: "Top 180 in CSE",
			Compensation:   "₹56,100 - ₹2,50,000+",
			SelectionRatio: "1 : 1,000",
			TotalMarks:     2025,
		},
	}
}

func createAICareer() models.Career {
	return models.Career{
		Name:        "AI & Machine Learning Engineer",
		Slug:        "ai-ml-engineer",
		Description: "Architect foundation models, fine-tune LLMs, design RAG pipelines, and deploy scalable inference infrastructure.",
		Category:    "tech",
		Difficulty:  "advanced",
		Tags:        []string{"technology", "ai", "ml", "deep learning", "data science"},
		Skills:      []string{"Python", "PyTorch", "TensorFlow", "Transformer Architectures", "Vector DBs", "CUDA", "Cloud Platforms"},
		ContentBlocks: []models.ContentBlock{
			{
				Type:   "hero",
				Title:  "AI & Machine Learning Engineer",
				Content: map[string]interface{}{
					"subtitle":   "Shape the future with Artificial Intelligence",
					"description": "AI Engineers build intelligent systems that can learn, reason, and assist in decision-making. This field is rapidly evolving with opportunities in research, product development, and deployment.",
				},
			},
			{
				Type:   "skills",
				Title:  "Technical Skills",
				Content: map[string]interface{}{
					"categories": []map[string]interface{}{
						{
							"name": "Core Programming",
							"skills": []string{"Python", "Mathematics", "Statistics", "Data Structures"},
						},
						{
							"name": "Machine Learning",
							"skills": []string{"Scikit-learn", "Model Training", "Hyperparameter Tuning", "Evaluation Metrics"},
						},
						{
							"name": "Deep Learning",
							"skills": []string{"PyTorch", "TensorFlow", "Neural Networks", "CNNs", "RNNs", "Transformers"},
						},
						{
							"name": "LLM & NLP",
							"skills": []string{"LangChain", "RAG", "Fine-tuning", "Prompt Engineering", "HuggingFace"},
						},
					},
				},
			},
			{
				Type:   "roadmap",
				Title:  "Learning Roadmap",
				Content: map[string]interface{}{
					"timeEstimate": "18-24 months",
					"steps": []models.RoadmapStep{
						{
							Title:           "Python & Math Basics",
							Description:     "Learn Python programming and brush up on linear algebra, calculus, and statistics",
							EstimatedMonths: 3,
							Status:          "completed",
						},
						{
							Title:           "Machine Learning Fundamentals",
							Description:     "Study ML algorithms and implement them from scratch",
							EstimatedMonths: 6,
							Status:          "in-progress",
						},
						{
							Title:           "Deep Learning",
							Description:     "Learn neural networks, CNNs, RNNs, and Transformers",
							EstimatedMonths: 8,
							Status:          "todo",
						},
						{
							Title:           "LLM & RAG",
							Description:     "Learn Large Language Models and Retrieval-Augmented Generation",
							EstimatedMonths: 12,
							Status:          "todo",
						},
						{
							Title:           "Deploy & Scale",
							Description:     "Learn model deployment, Docker, Kubernetes, and cloud platforms",
							EstimatedMonths: 18,
							Status:          "todo",
						},
					},
				},
			},
		},
		Roadmap: []models.RoadmapStep{
			{Title: "Python Programming", Description: "Master Python for data science", EstimatedMonths: 2},
			{Title: "ML Algorithms", Description: "Learn and implement ML algorithms", EstimatedMonths: 6},
			{Title: "Deep Learning", Description: "Study neural networks and deep learning", EstimatedMonths: 8},
			{Title: "LLM Projects", Description: "Build projects with LLMs and RAG", EstimatedMonths: 12},
		},
		Metrics: models.Metrics{
			DurationMonths: 18,
			EntryThreshold: "Bachelor's in CS/Related field",
			Compensation:   "$140k – $320k+ / ₹25L–70L+",
			SelectionRatio: "High demand",
			TotalMarks:     0,
		},
	}
}

func createFinanceCareer() models.Career {
	return models.Career{
		Name:        "Quantitative Trader",
		Slug:        "quantitative-trader",
		Description: "Develop mathematical models to identify market inefficiencies and execute algorithmic high-frequency trading strategies.",
		Category:    "finance",
		Difficulty:  "elite",
		Tags:        []string{"finance", "trading", "quant", "mathematics", "algorithmic"},
		Skills:      []string{"Stochastic Calculus", "C++/Rust", "Time Series Analysis", "Game Theory", "Python", "Machine Learning"},
		ContentBlocks: []models.ContentBlock{
			{
				Type:   "hero",
				Title:  "Quantitative Trader",
				Content: map[string]interface{}{
					"subtitle":   "Master the financial markets with mathematics",
					"description": "Quantitative traders use mathematical models, statistical analysis, and programming to identify and exploit market inefficiencies. This is one of the most competitive and rewarding careers in finance.",
				},
			},
			{
				Type:   "requirements",
				Title:  "Quantitative Finance Roadmap",
				Content: map[string]interface{}{
					"requirements": []map[string]interface{}{
						{
							"title":       "Mathematics",
							"description": "Master probability, statistics, stochastic calculus, and linear algebra",
							"icon":        "📐",
						},
						{
							"title":       "Programming",
							"description": "Expertise in C++, Python, and high-frequency trading systems",
							"icon":        "💻",
						},
						{
							"title":       "Finance",
							"description": "Deep understanding of markets, derivatives, and risk management",
							"icon":        "📈",
						},
					},
				},
			},
		},
		Roadmap: []models.RoadmapStep{
			{Title: "Mathematics Foundation", Description: "Learn probability and statistics", EstimatedMonths: 6},
			{Title: "Programming Mastery", Description: "Master C++ and Python for trading", EstimatedMonths: 8},
			{Title: "Financial Markets", Description: "Study market mechanics and instruments", EstimatedMonths: 6},
			{Title: "Quant Strategies", Description: "Develop and backtest trading strategies", EstimatedMonths: 12},
		},
		Metrics: models.Metrics{
			DurationMonths: 24,
			EntryThreshold: "Top quant programs / elite universities",
			Compensation:   "Elite base + performance bonus",
			SelectionRatio: "Very competitive",
			TotalMarks:     0,
		},
	}
}

func createDesignCareer() models.Career {
	return models.Career{
		Name:        "Product Designer",
		Slug:        "product-designer",
		Description: "Craft intuitive multi-platform user journeys, establish scalable design tokens and UI component libraries.",
		Category:    "design",
		Difficulty:  "intermediate",
		Tags:        []string{"design", "ui/ux", "product", "research", "visual design"},
		Skills:      []string{"Figma", "Interaction Design", "User Research", "Prototyping", "Design Systems", "Accessibility"},
		ContentBlocks: []models.ContentBlock{
			{
				Type:   "hero",
				Title:  "Product Designer",
				Content: map[string]interface{}{
					"subtitle":   "Design experiences that millions use daily",
					"description": "Product designers bridge the gap between user needs and business goals. They research, prototype, and ship designs that create meaningful user experiences across digital products.",
				},
			},
			{
				Type:   "skills",
				Title:  "Design Skills",
				Content: map[string]interface{}{
					"categories": []map[string]interface{}{
						{
							"name":   "UI/UX Design",
							"skills": []string{"Wireframing", "Prototyping", "User Testing", "Information Architecture"},
						},
						{
							"name":   "Tools",
							"skills": []string{"Figma", "Sketch", "Adobe Creative Suite", "Framer", "Webflow"},
						},
						{
							"name":   "Development",
							"skills": []string{"HTML/CSS", "Basic JavaScript", "Design Handoff", "Code Review"},
						},
					},
				},
			},
		},
		Roadmap: []models.RoadmapStep{
			{Title: "Design Fundamentals", Description: "Learn principles of visual design and typography", EstimatedMonths: 3},
			{Title: "UI/UX Practices", Description: "Master user research and wireframing", EstimatedMonths: 6},
			{Title: "Tool Mastery", Description: "Become expert in Figma and prototyping tools", EstimatedMonths: 4},
			{Title: "Portfolio Projects", Description: "Build real-world projects for your portfolio", EstimatedMonths: 6},
		},
		Metrics: models.Metrics{
			DurationMonths: 12,
			EntryThreshold: "Portfolio + Design degree or experience",
			Compensation:   "₹6L–25L+ depending on experience",
			SelectionRatio: "Good demand",
			TotalMarks:     0,
		},
	}
}
