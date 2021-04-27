use std::convert::TryInto;

pub fn nth_prime(nth: u64) -> i64 {
    let mut primes: Vec<i64> = Vec::new();
    let mut num = 2;
    loop {
        if is_prime(num) {
            primes.push(num);
            if primes.len() - 1 == nth.try_into().unwrap() {
                return num
            }
        }
        num += 1;
    }
}

fn is_prime(n: i64) -> bool {
    if n == 2 || n == 3 {
        return true;
    }
    let fln = n as f64;
    let upper = ((fln).sqrt()) as i64;
    for m in 2..upper+1 {
        if n % m == 0 {
            return false;
        }
    }
    return true;
}

#[test]
fn test_nth_prime(){
    assert_eq!(13, nth_prime(5));
}

#[test]
fn test_is_prime(){
    assert!(!is_prime(6));
    assert!(is_prime(7));
    assert!(!is_prime(12));
    assert!(is_prime(13));
}