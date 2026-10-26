classdef StateOfCharge
  properties
    % MaxCapacity is defined as kWh
    MaxCapacity (1,1) double
    % State, Min and Max Clamp are defined as percentage values
    State (1,1) double
    MinClamp (1,1) double = 0
    MaxClamp (1,1) double = 100
  end

  methods
    function soc = StateOfCharge(maxCapacity, initialState, lowLimit, highLimit)
      soc.MaxCapacity = maxCapacity;
      soc.State = initialState;
      soc.MinClamp = lowLimit;
      soc.MaxClamp = highLimit;

      soc = Charge(soc, 0);
    end

    function soc = Charge(soc, delta);
      soc.State = soc.State + (delta/soc.MaxCapacity) * 100;
      soc.State = min(max(soc.State, soc.MinClamp), soc.MaxClamp);
    end
  end
end

