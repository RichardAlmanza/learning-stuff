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

    function strs = Str(sm, strEnabled, strDisabled)
      strs = cell(1, length(sm.States));
      strs(sm.States == false) = { strDisabled };
      strs(sm.States == true) = { strEnabled };
    end
  end
end

