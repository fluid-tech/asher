package builder


const BasicTryCatch = `try {
    $error = 5/0;
}
catch ( DivideByZeroException $exception ) {
}
`
const BasicTryCatchFinally = `try {
    $error = 5/0;
}
catch ( DivideByZeroException $exception ) {
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