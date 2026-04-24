```
anton@sawdesk ~/p/versitygw (main)> export AWS_ENDPOINT_URL="http://127.0.0.1:7070"
anton@sawdesk ~/p/versitygw (main)> export AWS_ACCESS_KEY_ID='AKIAS37G7OCL2HF6XKO2'
anton@sawdesk ~/p/versitygw (main)> export AWS_SECRET_ACCESS_KEY="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"
anton@sawdesk ~/p/versitygw (main)> aws s3api list-buckets
{
    "Buckets": [
        {
            "Name": "objects",
            "CreationDate": "2026-04-24T18:00:40+00:00",
            "BucketRegion": "us-east-1"
        }
    ],
    "Owner": {
        "DisplayName": "",
        "ID": "AKIAS37G7OCL2HF6XKO2"
    },
    "Prefix": null
}
anton@sawdesk ~/p/versitygw (main)> aws s3api help
anton@sawdesk ~/p/versitygw (main)> aws s3api help
anton@sawdesk ~/p/versitygw (main)> aws s3api list-objects --bucket objects 
{
    "Contents": [
        {
            "Key": "Iskusstvennyy_intellekt_v_Rossii_2025.pdf",
            "LastModified": "2026-04-24T18:00:40+00:00",
            "ETag": "\"bc83dea1f6cd92a71bf4a9e5a8dd108e-4\"",
            "ChecksumAlgorithm": [
                "CRC64NVME"
            ],
            "ChecksumType": "FULL_OBJECT",
            "Size": 16824487,
            "StorageClass": "STANDARD",
            "Owner": {
                "ID": "AKIAS37G7OCL2HF6XKO2"
            }
        }
    ],
    "RequestCharged": null,
    "Prefix": null
}
anton@sawdesk ~/p/versitygw (main)> aws s3api get-object --bucket objects 'Iskusstvennyy_intellekt_v_Rossii_2025.pdf'

usage: aws [options] <command> <subcommand> [<subcommand> ...] [parameters]
To see help text, you can run:

  aws help
  aws <command> help
  aws <command> <subcommand> help

aws: error: the following arguments are required: --key

anton@sawdesk ~/p/versitygw (main) [252]> aws s3api get-object help
<string>:30: (WARNING/2) Inline literal start-string without end-string.
anton@sawdesk ~/p/versitygw (main)> aws s3api get-object --bucket objects --key Iskusstvennyy_intellekt_v_Rossii_2025.pdf ~/out.pdf 
{
    "AcceptRanges": "bytes",
    "LastModified": "2026-04-24T18:00:40+00:00",
    "ContentLength": 16824487,
    "ETag": "\"bc83dea1f6cd92a71bf4a9e5a8dd108e-4\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "ContentType": "application/pdf",
    "Metadata": {},
    "StorageClass": "STANDARD"
}
anton@sawdesk ~/p/versitygw (main)> aws s3api get-object --bucket objects --key Iskusstvennyy_intellekt_v_Rossii_2025.pdf ~/out.pdf
{
    "AcceptRanges": "bytes",
    "LastModified": "2026-04-24T18:00:40+00:00",
    "ContentLength": 16824487,
    "ETag": "\"bc83dea1f6cd92a71bf4a9e5a8dd108e-4\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "ContentType": "application/pdf",
    "Metadata": {},
    "StorageClass": "STANDARD"
}
anton@sawdesk ~/p/versitygw (main)> aws s3api get-object --bucket objects --key test.pdf --body ~/out.pdf

usage: aws [options] <command> <subcommand> [<subcommand> ...] [parameters]
To see help text, you can run:

  aws help
  aws <command> help
  aws <command> <subcommand> help

Unknown options: --body

anton@sawdesk ~/p/versitygw (main) [252]> aws s3api put-object --bucket objects --key test.pdf --body ~/out.pdf
{
    "ETag": "\"07305005c09fee770aa0ff6d7a31390e\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "Size": 16824487
}
anton@sawdesk ~/p/versitygw (main)> aws s3api put-object --bucket objects --key test.pdf --body ~/out.pdf
{
    "ETag": "\"07305005c09fee770aa0ff6d7a31390e\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "Size": 16824487
}
anton@sawdesk ~/p/versitygw (main)> aws s3api put-object --bucket objects --key test.pdf --body ~/out.pdf
{
    "ETag": "\"07305005c09fee770aa0ff6d7a31390e\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "Size": 16824487
}
anton@sawdesk ~/p/versitygw (main)> aws s3api put-object --bucket objects --key test.pdf --body ~/out.pdf
{
    "ETag": "\"07305005c09fee770aa0ff6d7a31390e\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "Size": 16824487
}
anton@sawdesk ~/p/versitygw (main)> aws s3api get-object --bucket objects --key test.pdf out.pdf
{
    "AcceptRanges": "bytes",
    "LastModified": "2026-04-24T18:09:28+00:00",
    "ContentLength": 16824487,
    "ETag": "\"07305005c09fee770aa0ff6d7a31390e\"",
    "ChecksumCRC64NVME": "o0TXEmuTdGo=",
    "ChecksumType": "FULL_OBJECT",
    "ContentType": "binary/octet-stream",
    "Metadata": {},
    "StorageClass": "STANDARD"
}

```