# API Payload Samples

This folder contains JSON payloads for testing the Career Architecture API.

## Files

### Career Creation Payloads

1. **create-career.json** - Complete UPSC IAS career with all content blocks
2. **create-career-ai.json** - AI & Machine Learning Engineer career
3. **create-career-finance.json** - Quantitative Trader career
4. **create-career-design.json** - Product Designer career
5. **create-career-ies.json** - Complete UPSC Engineering Services (IES/ESE) architectural blueprint

### Update Payload

6. **update-career.json** - Example update request

### Bulk Seed

7. **seed-all-careers.json** - All sample careers for bulk seeding

## API Usage Examples

### Create a new career

```bash
curl -X POST http://localhost:8080/api/careers \
  -H "Content-Type: application/json" \
  -d @payloads/create-career.json
```

### Create AI Engineer career

```bash
curl -X POST http://localhost:8080/api/careers \
  -H "Content-Type: application/json" \
  -d @payloads/create-career-ai.json
```

### Update a career (replace ID with actual career ID)

```bash
curl -X PUT http://localhost:8080/api/careers/{career-id} \
  -H "Content-Type: application/json" \
  -d @payloads/update-career.json
```

### Get all categories

```bash
curl http://localhost:8080/api/categories
```

### Get all careers

```bash
curl http://localhost:8080/api/careers
```

### Get career by slug

```bash
curl http://localhost:8080/api/careers/slug/upsc-ias
```

### Filter by category

```bash
curl "http://localhost:8080/api/careers?category=tech"
```

## Content Block Types

- `hero` - Hero section with title, subtitle, description
- `requirements` - Requirements checklist with icons
- `skills` - Skills categories with tags
- `assessment` - Assessment stages (exams, interviews)
- `roadmap` - Detailed roadmap with steps
- `faq` - Frequently asked questions
- `resources` - Learning resources with links
- `metrics` - Key career metrics

## Roadmap Step Fields

```json
{
  "title": "Step Title",
  "description": "Step description",
  "estimated_months": 3,
  "dependencies": ["step1", "step2"],
  "resources": [
    {
      "title": "Resource Title",
      "url": "https://example.com",
      "type": "video"
    }
  ],
  "status": "todo"
}
```

## Response Format

All API responses follow this format:

```json
{
  "success": true,
  "message": "Career fetched successfully",
  "data": {...}
}
```
