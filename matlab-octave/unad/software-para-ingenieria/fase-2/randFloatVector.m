% randFloatVector returns a vector with random float numbers
% within range (min, max) and the dimension is (1, size)

function newVect = randFloatVector(min, max, size)
  newVect = min + (max - min) .* rand(1, size);
end

