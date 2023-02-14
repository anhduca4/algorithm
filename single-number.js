
/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
  console.log(nums);
  let xor = 0;
  for (const num of nums) {
    xor ^= num;
  }
  return xor;
};

console.log(singleNumber([2, 2, 1]));
console.log(singleNumber([4, 1, 2, 1, 2]));
