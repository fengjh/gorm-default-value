# Gorm Default Value

Gorm takes the default value back from database after you create a new record with no value specified.

## Quick Started 

    # Get code
    $ go get -u github.com/fengjh/gorm-default-value

    # Set environment variables
    $ source .env

    # Setup postgres database
    $ postgres=# CREATE USER gorm_default_value;
    $ postgres=# CREATE DATABASE gorm_default_value OWNER gorm_default_value;

    # Run Application
    $ go test -v

## Test Result

    === RUN   TestCreateProductPriceWithDefaultValue
    BeforeCreate --> 0
    AfterCreate --> 0.01
    BeforeCreateTransactionCommit --> 0.01
    AfterCreateTransactionCommit --> 0.01
    --- PASS: TestCreateProductPriceWithDefaultValue (0.00s)
    === RUN   TestSaveProductPriceWithDefaultValue
    BeforeCreate --> 0
    AfterCreate --> 0.01
    BeforeCreateTransactionCommit --> 0.01
    AfterCreateTransactionCommit --> 0.01
    --- PASS: TestSaveProductPriceWithDefaultValue (0.00s)
    PASS
    ok  	github.com/fengjh/gorm-default-value	0.072s
