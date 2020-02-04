from pyspark.sql.functions import *
from pyspark.context import SparkContext
from awsglue.transforms import *
from awsglue.context import GlueContext
from awsglue.job import Job

glue_context = GlueContext(SparkContext.getOrCreate())
spark = glue_context.spark_session
job = Job(glue_context)
job.init('ddb2s3')

dyf = glue_context.create_dynamic_frame.from_options(
    connection_type='dynamodb',
    connection_options={
        'dynamodb.input.tableName': 'NatureRemo',
        'dynamodb.splits': '72' # Should be (DPU-1)*8
    }
)

dyf.printSchema()
dyf = dyf.resolveChoice(specs = [('value','cast:double')])
dyf = dyf.repartition(1) # You need to modify this number to your preferred file number.

glue_context.write_dynamic_frame_from_options(
    frame=dyf,
    connection_type='s3',
    connection_options={
        'path': "s3://example-bucket/glue/ddb2s3/"
    },
    format='parquet'
)

# You will need to specify partitionKeys if you want to write data with hive partitioning based on column 'year', 'month', and 'day
#glue_context.write_dynamic_frame_from_options(
#    frame=dyf,
#    connection_type='s3',
#    connection_options={
#        'path': "s3://example-bucket/glue/ddb2s3/"
#        'partitionKeys': ["year", "month", "day"]
#    },
#    format='parquet'
#)

job.commit()