# Medical API System

A simple REST API for managing patient medical records, examinations, samples, prescriptions, and referrals.

## Core Processes

### 1. Patient Management

- Register new patients with basic information
- View patient details
- Update patient information
- Remove patients from the system

### 2. Examination Process

- Create medical examinations for patients
- Record anamnesis (patient history) and diagnosis
- View examination details
- Filter examinations by patient

### 3. Sample Collection & Evaluation

**Workflow:**
1. Create a sample for an examination (blood, urine, tissue, etc.)
2. System automatically evaluates the sample
3. Results are recorded with the sample
4. Based on evaluation, a prescription is automatically generated

### 4. Prescription Management

**Workflow:**
1. Prescriptions can be created manually or automatically from sample evaluation
2. All prescriptions require validation by a doctor
3. Validated prescriptions can be sent to pharmacy
4. Prescriptions track validation and delivery status

### 5. Referral System

- Create referrals to specialists
- Associate referrals with specific examinations
- Track referral reasons and specialist details

## API Usage

The API runs on port 8080 and all endpoints are accessible under `/api` prefix.

### Example Flow:

1. Create a patient (`POST /api/patients`)
2. Create an examination (`POST /api/examinations`)
3. Add sample to examination (`POST /api/samples`)
4. System automatically evaluates sample and generates prescription
5. Validate prescription (`POST /api/prescriptions/{id}/validate`)
6. Send prescription to pharmacy (`POST /api/prescriptions/{id}/send`)
7. Create referral if needed (`POST /api/referrals`)

See OpenAPI documentation for detailed endpoint specifications.