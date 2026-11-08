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

    function soc = Charge(soc, delta)
      soc.State = soc.State + (delta/soc.MaxCapacity) * 100;
      soc.State = min(max(soc.State, soc.MinClamp), soc.MaxClamp);
    end

    function [soc, records] = ChargeRecords(soc, deltaRecords)
      len = length(deltaRecords);
      records = zeros(1, len);

      soc = Charge(soc, deltaRecords(1));
      records(1) = soc.State;

      for index = 2:len
          soc = Charge(soc, deltaRecords(index));
          records(index) = soc.State;
      end
    end

  end
end

