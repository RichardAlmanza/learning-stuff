% generatePowerGenerationData returns two vectors
%   solar vector is concave down parabola, which reach its peaks
%     at 12 o'clock, the curve goes from 5 to 19, both ends are 0.
%     function described is y = -(25/49) * (x - 12)^2 + 25
%   wind vector have random generated values within range (0,15)

function [solar, wind] = generatePowerGenerationData()
  hours = 1:24;

  % Solar energy definition
  solar = zeros(1, 24);
  curve = -(25/49) .* (hours - 12).^2 + 25;
  solar(5:19) = curve(curve >= 0);

  % Wind energy definition
  wind = 15 .* rand(1, 24);
end

