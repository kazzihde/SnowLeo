# ❄︎ SnowLeo ❄︎

A self-hosted OSINT platform for collecting, analyzing, and organizing publicly available intelligence.

## Features (Planned)

### 〉Core Intelligence Modules

#### Domain Intelligence
- WHOIS lookup (registrar, creation/expiry, name servers)
- Domain history tracking
- Registrar abuse/contact metadata

#### DNS Analysis
- A / AAAA / MX / TXT record resolution
- Passive DNS enrichment
- GeoIP mapping of resolved IPs
- CDN detection (e.g. Cloudflare, Akamai)

#### Website Analysis
- HTTP header inspection
- Security misconfiguration detection
- Technology fingerprinting (CMS/framework detection)
- Headless screenshot capture

#### Certificate Transparency Monitoring
- CT log querying
- Subdomain extraction from certificates
- Expired/renewed certificate tracking
- Certificate chain analysis

### 〉Discovery & Enrichment

#### Subdomain Discovery
- Passive CT-based enumeration
- Optional DNS-based discovery
- Aggregated domain surface mapping

#### OSINT Enrichment Engine
- Automatic expansion of input entities (domain/IP/email)
- Cross-linking between datasets (DNS ↔ WHOIS ↔ certificates)
- Confidence-based relationship mapping

#### Footprint Analyzer
- Relationship graph generation between entities
- IP ↔ domain ↔ organization linking
- Investigation timeline tracking
- Graph-based visualization (future)

### 〉Media & Metadata Analysis

#### Reverse Image Search
- Multi-source lookup (search engine integrations)
- Image matching across public indexes
- Metadata correlation (when available)

#### EXIF Metadata Extractor
- Camera/device metadata extraction
- GPS coordinate parsing (if present)
- Timestamp validation and inconsistencies

#### Media Intelligence (Future)
- Optional face clustering on user-provided datasets
- Strictly non-identifying grouping of similar images
- No person identification or scraping of individuals

### 〉Reporting & Workflow

#### Search History
- Query tracking per session
- Investigation timeline view
- Session-based grouping

#### Report Exporting
- Markdown reports (initial)
- JSON export for API usage
- PDF report generation (future)
- Structured “case file” output

#### Investigation Workspaces (Future)
- Multi-query projects
- Saved entities and findings
- Timeline and graph-based investigation view

### 〉Security & Access (Planned)
- Authentication system (API keys / user accounts)
- Rate limiting per user
- Request logging and usage tracking

### 〉Advanced Intelligence (Future)
- Passive OSINT data collection layer
- Historical dataset tracking (DNS/WHOIS changes over time)
- Intelligence graph database backend
- Automated enrichment pipelines
- Optional AI-generated summaries of findings

## ﹗Tech Stack﹗

### Frontend: React
- Dashboard for queries, results, and investigation workflows

### Backend: Go
- Core API, routing, and high-performance OSINT queries

### Data Processing / Automation: Python
- Enrichment scripts, parsing, and external data integrations

### Database: MariaDB
- Stores structured intelligence data and relationships

### Optional Extensions: Node.js / Go / Python workers
- Real-time updates, background jobs, and specialized modules

## Status

🚧 IN ACTIVE DEVELOPMENT
