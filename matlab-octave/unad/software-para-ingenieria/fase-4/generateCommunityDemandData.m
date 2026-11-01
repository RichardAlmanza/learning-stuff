% generateCommunityDemandData returns one vector
%   this vector follows these rules:
%   - [1, 5] U [22, 24] indices have values within range (2, 4)
%   - [6, 17] indices have values within range (6, 10)
%   - [18, 21] indices have the 15 number

function communityDemand = generateCommunityDemandData()
  % Apply minimum demand (2 ~ 4)
  communityDemand = randFloatVector(2, 4, 24);

  % Morning demand (6 ~ 10)
  communityDemand(6:17) = randFloatVector(6, 10, 12);

  % Peak demand (15)
  communityDemand(18:21) = 15 .* ones(1,4);
end
