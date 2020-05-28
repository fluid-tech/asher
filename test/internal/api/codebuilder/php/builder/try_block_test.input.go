package builder

const BasicTryCatch = `try {
    $error = 5/0;
}
catch ( DivideByZeroException $exception ) {
    echo "There was an exception";
}
`
const BasicTryCatchFinally = `try {
    $error = 5/0;
}
catch ( DivideByZeroException $exception ) {
    echo "There was an exception";
}
finally {
    echo "Hello I am in finally";
}
`
const BasicTryFinally = `try {
    $error = 5/0;
}
finally {
    echo "Hello I am in finally";
}
`

const TryCatchMultipleStatement = `try {
    $error = 5/0;
    echo "Hello in try block";
}
catch ( DivideByZeroException $exception ) {
    echo "There was an exception";
    return $exception;
}
`

const TryCatchFinallyMultipleStatement = `try {
    $error = 5/0;
    echo "Hello in try block";
}
catch ( DivideByZeroException $exception ) {
    echo "There was an exception";
    return $exception;
}
finally {
    echo "Hello I am in finally";
    echo "This block gets execute everytime";
}
`
const TryFinallyMultipleStatement = `try {
    $error = 5/0;
    echo "Hello in try block";
}
finally {
    echo "Hello I am in finally";
    echo "This block gets execute everytime";
}
`
