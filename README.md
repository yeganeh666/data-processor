# Data Processor - Input Component

This repository contains the implementation of the Input Component for the Data Processor system. The input component is responsible for handling input data efficiently and ensuring uniqueness based on unique identifiers.

## Input-Process-Storage-Output (IPSO) Concept

<p align="center">
  <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTFpwG7pCbhOA8tbZG_RzNqbR8Xb4YmqI7Pcv5UqbrMoxTjCd6NLqQhEATGQxgetDNf7m8&usqp=CAU" alt="Alt Text">
</p>

The Data Processor project aligns with the Input-Process-Storage-Output (IPSO) concept, which expands upon the traditional Input-Process-Output (IPO) model by incorporating a dedicated storage component. The IPSO model emphasizes the importance of securely storing processed data for future retrieval and utilization.

In our Data Processor project, we can observe the following relationship with the IPSO model:

**Input**: The system receives input data, consisting of unique identifiers and user identifiers, from various sources using popular protocols like HTTP.

**Process**: The input data undergoes processing stages, which include ensuring uniqueness, preventing duplicates, and enforcing user-specific rate limits. These steps guarantee the integrity and quality of the processed data.

**Storage**: Although the storage component itself is not implemented within this specific project, it is essential to acknowledge its significance. The processed data is directed towards a designated storage space, ensuring its preservation and availability for future access and analysis.

**Output**: The final output of our system encompasses not only the processed data but also its seamless integration into a storage infrastructure. This enables efficient retrieval, long-term storage, and utilization of the data for downstream applications or analysis.

## Requirements

The input component is designed to meet the following requirements:

1. **Unique Data Handling**: The system ensures that each data entry is associated with a globally unique identifier (GUID) and a user identifier. It prevents the receipt and processing of duplicate data as much as possible.

2. **User Quotas**: Each user has specific quotas for the number of requests they can make per minute and the total data volume they can send per month.

3. **Quota Enforcement**: If a user exceeds their allowed quotas, the system blocks further data reception from that user. Data deletion after receipt is not allowed due to quota violation.

4. **Response Time Commitment**: The system is committed to storing input data within a specific timeframe after its submission.

## Getting Started

1. **Clone the repository**:

```
git clone https://github.com/yeganeh666/data_processor.git
```

2. **Build the Docker containers**:

```
make docker
```

3. **Access the web service at**: `http://localhost:8080`.

## API Documentation

API documentation is available in Swagger format. You can view the Swagger docs by visiting http://localhost:8080/swagger/index.html in your web browser.
