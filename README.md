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

    === RUN   TestCreateProductWithNoValueSpecified
    BeforeCreate --> 0
    BeforeCreate -->
    AfterCreate --> 0.01
    AfterCreate --> clothing
    BeforeCreateTransactionCommit --> 0.01
    BeforeCreateTransactionCommit --> clothing
    AfterCreateTransactionCommit --> 0.01
    AfterCreateTransactionCommit --> clothing
    --- PASS: TestCreateProductWithNoValueSpecified (0.00s)
    === RUN   TestSaveProductPriceWithNoValueSpecified
    BeforeCreate --> 0
    BeforeCreate -->
    AfterCreate --> 0.01
    AfterCreate --> clothing
    BeforeCreateTransactionCommit --> 0.01
    BeforeCreateTransactionCommit --> clothing
    AfterCreateTransactionCommit --> 0.01
    AfterCreateTransactionCommit --> clothing
    --- PASS: TestSaveProductPriceWithNoValueSpecified (0.00s)
    === RUN   TestCreateProductPriceWithZeroValueSpecified
    BeforeCreate --> 0
    BeforeCreate -->
    AfterCreate --> 0.01
    AfterCreate --> clothing
    BeforeCreateTransactionCommit --> 0.01
    BeforeCreateTransactionCommit --> clothing
    AfterCreateTransactionCommit --> 0.01
    AfterCreateTransactionCommit --> clothing
    --- PASS: TestCreateProductPriceWithZeroValueSpecified (0.00s)
    PASS
    ok  	github.com/fengjh/gorm-default-value	0.075s
