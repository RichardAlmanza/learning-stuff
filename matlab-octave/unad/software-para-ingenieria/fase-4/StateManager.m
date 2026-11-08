classdef StateManager
  properties
    Names (1,:)
    Thresholds (1,:) double
    States (1,:) logical = true
  end

  methods
    function sm = StateManager(names, thresholds)
      sm.Names = names;
      sm.Thresholds = thresholds;
    end

    function sm = Evaluate(sm, values)
      sm.States = values > sm.Thresholds;
    end

    function [sm, records] = EvaluateRecords(sm, values)
      statesLen = length(sm.Thresholds);
      len = length(values);
      records = zeros(len, statesLen);

      sm = sm.Evaluate(values(1));
      records(1, :) =  sm.States;

      for index = 2:len
          sm = sm.Evaluate(values(index));
          records(index, :) = sm.States;
      end
    end

    function strs = Str(sm, strEnabled, strDisabled)
      strs = cell(1, length(sm.States));
      strs(sm.States == false) = { strDisabled };
      strs(sm.States == true) = { strEnabled };
    end
  end
end

